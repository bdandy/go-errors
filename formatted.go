package typed_errors

import "fmt"

// formattedError represents error with additional params for Sprintf format
type formattedError struct {
	typedError
	args []interface{}
}

// Error implements errors interface
func (e formattedError) Error() string {
	return e.String()
}

// ErrorString() implements stringer interface
func (e formattedError) String() string {
	return fmt.Sprintf(e.typedError.Error(), e.args...)
}

// Wrap adds cause to the formattedError error and return wrapped
func (e formattedError) Wrap(cause error) wrapped {
	return Wrap(e, cause)
}

// WithStack implements TypedError
func (e formattedError) WithStack() wrapped {
	return WrapWithStack(e, nil)
}
