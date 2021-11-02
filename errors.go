package typed_errors

import (
	"fmt"
)

type err interface {
	error
	Is(err error) bool
}

// formatted represents error with additional params for Sprintf format
type formatted struct {
	err
	args []interface{}
}

// Error implements errors interface
func (e formatted) Error() string {
	return fmt.Sprintf(e.err.Error(), e.args...)
}

// cause is custom error type for better error handling
type cause struct {
	err
	cause error
}

// Unwrap implements errors.Unwrap interface
func (e cause) Unwrap() error {
	return e.cause
}

// Error implements error interface; prints both error and cause
func (e cause) Error() string {
	if e.cause == nil {
		return e.err.Error()
	}

	return fmt.Sprintf("%s %#v", e.err.Error(), e.cause)
}
