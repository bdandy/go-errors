package typed_errors

// String which used in const as error
type String string

// New returns new typed error
func (e String) New() typedError {
	return typedError{e}
}

// NewWithArgs returns new error which would be formattedError
// Note: if two errors has different arguments they still would be equal
// Used Sprintf for formatting, so use #Wrap method instead of %w for wrapping
func (e String) NewWithArgs(args ...interface{}) formattedError {
	return formattedError{
		typedError: e.New(),
		args:       args,
	}
}

// Error implements errors interface
func (e String) Error() string {
	return e.String()
}

// typedError implements stringer interface
func (e String) String() string {
	return string(e)
}
