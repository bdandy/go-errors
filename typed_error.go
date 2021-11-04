package typed_errors

type TypedError interface {
	error

	Is(err error) bool
	WithArgs(args ...interface{}) formattedString
	Wrap(cause error) wrapped
}
