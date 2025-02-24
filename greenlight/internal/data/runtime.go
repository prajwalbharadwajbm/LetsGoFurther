package data

import (
	"fmt"
	"strconv"
)

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
