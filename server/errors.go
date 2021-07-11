package server

import "errors"

var (
	// ErrJSONMalformed is returned if the JSON contains malformed input
	ErrJSONMalformed = errors.New("request body contains malformed JSON")
)
