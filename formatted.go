package typed_errors

import "fmt"

// formattedString represents error with additional params for Sprintf format
type formattedString struct {
	ErrorString
	args []interface{}
}

// Error implements errors interface
func (e formattedString) Error() string {
	return e.String()
}

// ErrorString() implements stringer interface
func (e formattedString) String() string {
	return fmt.Sprintf(e.ErrorString.Error(), e.args...)
}

// Wrap adds cause to the formattedString error and return wrapped
func (e formattedString) Wrap(cause error) wrapped {
	return Wrap(e, cause)
}

// Stack implements TypedError
// NOOP as formattedString isn't wrapped
func (e formattedString) Stack() string {
	return ""
}
