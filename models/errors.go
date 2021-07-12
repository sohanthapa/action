package models

import "errors"

var (
	// ErrActionFieldEmpty is returned when action field is empty value
	ErrActionFieldEmpty = errors.New("action field is empty")
	// ErrTimeFieldZero is returned when time field is zero value
	ErrTimeFieldZero = errors.New("time field is zero value")
)
