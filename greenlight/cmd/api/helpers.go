package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Encode the data to JSON, bubble up the error if exists.
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')

	// Looping through the header map and adding each header to the http.ResponseWriter header map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Set Response Headers to application/json and utf-8 encoding, as this application will be interacting with public internet.
	// And according to RFC standards, we "must" set utf-8 encoding when interacting with json on any open-systems.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("charset", "utf-8")

	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// Use http.MaxBytesReader() to limit the size of the request body to 1MB.
	maxBytes := 1_048_567
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Initialize json.Decoder, and call DisallowUnkownFields() method on it before decoding.
	// Is the JSON input has unknown fields which can't be mapped to target destination, it returns
	// error instead of ignoring it.
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		// must use a non-nil pointer for Decode() i.e., dst, else we find json.InvalidUnmarshalError error at runtime.
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Offset == 1 {
				return errors.New("body must contain a JSON object")
			}
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	// Additionally to check if there is more than one object being sent in body,
	// We can check by calling decode the second time comapring it with EOF, If not,
	// Then we can error out saying body contain single JSON Value.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON Value")
	}
	return nil
}
