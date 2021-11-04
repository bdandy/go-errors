package typed_errors

type TypedError interface {
	error

	Is(err error) bool        // Is checks if TypedError and provided error has same type
	Wrap(cause error) wrapped // Wrap wraps provided error
	WithStack() wrapped       // Stack stores call stack and returns wrapped error
}

type ErrorString = String

type typedError struct {
	ErrorString
}

// Wrap adds cause to the String error and return wrapped
func (e typedError) Wrap(cause error) wrapped {
	return Wrap(e, cause)
}

// WithStack implements TypedError interface
func (e typedError) WithStack() wrapped {
	return WrapWithStack(e, nil)
}

// Is implements errors.Is
func (e typedError) Is(err error) bool {
	ee, ok := err.(String)
	if !ok {
		return false
	}

	return e.ErrorString == ee
}
