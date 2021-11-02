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

// wrapped is custom error type for better error handling
type wrapped struct {
	err
	cause error
}

// Unwrap implements errors.Unwrap interface
func (e wrapped) Unwrap() error {
	return e.cause
}

// Error implements error interface; prints both error and cause
func (e wrapped) Error() string {
	if e.cause == nil {
		return e.err.Error()
	}

	return fmt.Sprintf("%s %#v", e.err.Error(), e.cause)
}
