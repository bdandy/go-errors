package typed_errors

import (
	"fmt"
)

type TypedError interface {
	error
	Is(err error) bool
	WithArgs(args ...interface{}) formatted
	Wrap(err error) wrapped
}

// formatted represents error with additional params for Sprintf format
type formatted struct {
	TypedError
	args []interface{}
}

// Error implements errors interface
func (e formatted) Error() string {
	return fmt.Sprintf(e.TypedError.Error(), e.args...)
}

// wrapped is custom error type for better error handling
type wrapped struct {
	TypedError
	cause error
}

// Unwrap implements errors.Unwrap interface
func (e wrapped) Unwrap() error {
	return e.cause
}

// Error implements error interface; prints both error and cause
func (e wrapped) Error() string {
	if e.cause == nil {
		return e.TypedError.Error()
	}

	return fmt.Sprintf("%s %#v", e.TypedError.Error(), e.cause)
}
