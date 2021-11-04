package typed_errors

import (
	"errors"
)

// String is alias to ErrorString
type String = ErrorString

// ErrorString implements TypedError
type ErrorString string

// Is implements errors.Is
func (e ErrorString) Is(err error) bool {
	var ee TypedError
	if !errors.As(err, &ee) {
		return false
	}

	return e == ee
}

// Error implements errors interface
func (e ErrorString) Error() string {
	return e.String()
}

// ErrorString implements stringer interface
func (e ErrorString) String() string {
	return string(e)
}

// WithArgs returns new error which would be formattedString
// Note: args are not impact on errors.Is; multiple call WithArgs sets most recent ones
// so if two errors has different arguments they still would be equal
// Use Sprintf for formatting, so use #Wrap method instead of %w for wrapping
func (e ErrorString) WithArgs(args ...interface{}) formattedString {
	return formattedString{
		ErrorString: e,
		args:        args,
	}
}

// Wrap adds cause to the ErrorString error and return wrapped
func (e ErrorString) Wrap(cause error) wrapped {
	return Wrap(e, cause)
}
