package typed_errors

import (
	"strings"
)

type Wrapped interface {
	Unwrap() error
}

// wrapped is custom error type for better error handling
type wrapped struct {
	TypedError
	cause error
}

// Wrap adds cause to the error and return new wrapped error
func (e wrapped) Wrap(cause error) wrapped {
	return Wrap(e, cause)
}

// Unwrap implements errors.Unwrap interface
func (e wrapped) Unwrap() error {
	return e.cause
}

// Error implements error interface; prints both error and cause
func (e wrapped) Error() string {
	var builder strings.Builder
	builder.WriteString(e.TypedError.Error())

	if e.cause != nil {
		const separator = ": "
		builder.WriteString(separator)
		builder.WriteString(e.cause.Error())
	}

	return builder.String()
}

// Wrap wraps an error with cause
func Wrap(err, cause error) (e wrapped) {
	te, ok := err.(TypedError)
	if !ok {
		e.TypedError = ErrorString(err.Error())
		e.cause = cause
		return
	}

	e.TypedError = te
	e.cause = cause

	return
}
