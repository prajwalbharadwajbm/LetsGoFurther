package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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
