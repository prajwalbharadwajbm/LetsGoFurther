package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

// MarshalJSON implements the json.Marshaler interface for the Runtime type.
// This allows us to customize how Runtime values are encoded to JSON.
// It converts the Runtime value (in minutes) to a string in the format "X mins"
// and returns the JSON-encoded string representation.
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}

// Implement UnmarshalJSON() Method on the runtime type so that it satisfies
// the json.Unmarshaler interface.
// We need to use pointer receiver for this to work, Otherwise, we will only be
// modifying a copy (which is discards when the method returns)
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// It is expected to receive qouted string value.
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// Split the string.
	parts := strings.Split(unquotedJSONValue, " ")

	// Sanity check if 2 parts of string exists after split and second part has mins in it.
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// Convert string to int32
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// deference and convert int32 to runtime format.
	*r = Runtime(i)
	return nil
}
