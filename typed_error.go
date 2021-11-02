package typed_errors

import (
	"errors"
	"fmt"
)

type TypedError interface {
	error
	fmt.Stringer

	Is(err error) bool
	WithArgs(args ...interface{}) formatted
	Wrap(err error) wrapped
}

// String implements TypedError
type String string

// Is implements errors.Is
func (e String) Is(err error) bool {
	var ee String
	if !errors.As(err, &ee) {
		return false
	}

	return e == ee
}

// Error implements errors interface
func (e String) Error() string {
	return string(e)
}

// String implements stringer interface
func (e String) String() string {
	return e.Error()
}

// Wrap adds cause to the String error and return wrapped
func (e String) Wrap(err error) wrapped {
	return wrapped{
		TypedError: e,
		cause:      err,
	}
}

// WithArgs returns new error which would be formatted
// Note: args are not impact on errors.Is; multiple call WithArgs sets most recent ones
// so if two errors has different arguments they still would be equal
// Use Sprintf for formatting, so use #Wrap method instead of %w for wrapping
func (e String) WithArgs(args ...interface{}) formatted {
	return formatted{
		TypedError: e,
		args:       args,
	}
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

// String() implements stringer interface
func (e formatted) String() string {
	return e.Error()
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

	return fmt.Sprintf("%s %s", e.TypedError.Error(), e.cause.Error())
}
