package typed_errors

import (
	"errors"
	"fmt"
)

type TypedError interface {
	error
	fmt.Stringer

	Is(err error) bool
	WithArgs(args ...interface{}) TypedError
	Wrap(err error) TypedError
}

// String implements TypedError
type String string

// Is implements errors.Is
func (e String) Is(err error) bool {
	var ee TypedError
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
func (e String) Wrap(cause error) TypedError {
	return Wrap(e, cause)
}

// WithArgs returns new error which would be formatted
// Note: args are not impact on errors.Is; multiple call WithArgs sets most recent ones
// so if two errors has different arguments they still would be equal
// Use Sprintf for formatting, so use #Wrap method instead of %w for wrapping
func (e String) WithArgs(args ...interface{}) TypedError {
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

// Wrap adds cause to the error and return new wrapped error
func (e wrapped) Wrap(cause error) TypedError {
	return Wrap(e, cause)
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

	return fmt.Sprint(e.TypedError.Error(), ": ", e.cause.Error())
}

// Wrap returns wrapped error
func Wrap(err error, cause error) (te TypedError) {
	var ok bool
	if te, ok = err.(TypedError); !ok {
		te = String(err.Error())
	}

	return wrapped{
		TypedError: te,
		cause:      cause,
	}
}
