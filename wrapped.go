package typed_errors

import (
	"bytes"
	"runtime"
	"strconv"
)

type Wrapped interface {
	Unwrap() error
}

// wrapped is custom error type for better error handling
type wrapped struct {
	TypedError
	cause     error
	stackFunc func() string
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
	var builder bytes.Buffer
	builder.WriteString(e.TypedError.Error())

	const separator = ": "
	if e.cause != nil {
		builder.WriteString(separator)
		builder.WriteString(e.cause.Error())
	}

	return builder.String()
}

func (e wrapped) withStack() wrapped {
	if e.stackFunc == nil {
		var callers [1 << 5]uintptr

		runtime.Callers(4, callers[:])
		e.stackFunc = func() string {
			frames := runtime.CallersFrames(callers[:])

			var buf bytes.Buffer
			for {
				frame, more := frames.Next()
				buf.WriteString(frame.Function)
				buf.WriteString("\n\t")
				buf.WriteString(frame.File)
				buf.WriteString(":")
				buf.WriteString(strconv.Itoa(frame.Line))
				buf.WriteString("\n")
				if !more {
					break
				}
			}

			return buf.String()
		}
	}

	return e
}

// Stack implements TypedError
// Stack returns original stack which was placed to wrapped error
func (e wrapped) Stack() string {
	if e.stackFunc != nil {
		return e.stackFunc()
	}

	return ""
}

// Wrap wraps an error, adds its cause and stack trace
func Wrap(err, cause error) (e wrapped) {
	te, ok := err.(TypedError)
	if !ok {
		e.TypedError = String(err.Error()).New()
		e.cause = cause
		return e
	}

	e.TypedError = te
	e.cause = cause

	return e
}

// WrapWithStack wraps cause with err and stores stack trace
func WrapWithStack(err, cause error) wrapped {
	return Wrap(err, cause).withStack()
}

// Stack returns stack if err is wrapped or implements Stack() string method
func Stack(err error) string {
	we, ok := err.(interface{ Stack() string })
	if !ok {
		return ""
	}

	return we.Stack()
}
