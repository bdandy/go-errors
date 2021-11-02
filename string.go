package typed_errors

import "errors"

// String is simple error type
type String string

// Is implements errors.Is
func (e String) Is(err error) bool {
	var ee String
	if !errors.As(err, &ee) {
		return false
	}

	return e == ee
}

// wrapped implements error interface
func (e String) Error() string {
	return string(e)
}

// Wrap adds cause to the String error and return wrapped
func (e String) Wrap(err error) wrapped {
	return wrapped{
		err:   e,
		cause: err,
	}
}

// WithArgs returns new error which would be formatted
// Note: args are not impact on errors.Is,
// so if two errors has different arguments they still would be equal
func (e String) WithArgs(args ...interface{}) formatted {
	return formatted{
		err:  e,
		args: args,
	}
}
