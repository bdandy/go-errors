package errors

import (
	"errors"
	"fmt"
	"testing"
)

const (
	err    = String("err")
	errFmt = String("err %s")
)

func TestTypedError_Is(t *testing.T) {
	type args struct {
		err error
	}

	var (
		cause   = errors.New("cause")
		wrapped = err.New().Wrap(cause)
		nested  = errFmt.New("test").Wrap(wrapped)
	)

	tests := []struct {
		name string
		e    error
		args args
		want bool
	}{
		{"same error", err.New(), args{err}, true},
		{"same error with args", errFmt.New("formattedError"), args{errFmt}, true},
		{"error has different type but same text", err.New(), args{errors.New("err")}, false},
		{"wrapped error", wrapped, args{err}, true},
		{"wrapped error check cause", wrapped, args{cause}, true},
		{"nested wrapped error", nested, args{cause}, true},
		{"nested wrapped error", nested, args{err}, true},
		{"nested wrapped error", nested, args{errors.New("cause")}, false},
		{"different errors", err.New(), args{String("other")}, false},
		{"error was cause", wrapped, args{fmt.Errorf("caused by %w", err)}, false},
		{"error is nil", err.New(), args{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Is(tt.e, tt.args.err); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Wrap(t *testing.T) {
	type args struct {
		err   Error
		cause error
	}

	var (
		cause = errors.New("root")
	)

	tests := []struct {
		name string
		args args
		want string
	}{
		{"error", args{err.New(), cause}, "err: root"},
		{"error with args", args{errFmt.New("formattedError"), cause}, "err formattedError: root"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.err.Wrap(tt.args.cause); got.Error() != tt.want {
				t.Errorf("Wrap() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}

func TestTypedError_Unwrap(t *testing.T) {
	var cause = errors.New("cause")

	tests := []struct {
		name string
		e    wrapped
		want error
	}{
		{"without cause", err.New().Wrap(nil), nil},
		{"with cause", err.New().Wrap(cause), cause},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Unwrap(); got != tt.want {
				t.Errorf("Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypedError_As(t *testing.T) {
	var (
		chain    = fmt.Errorf("second: %w", fmt.Errorf("first: %w", err.New()))
		chainFmt = fmt.Errorf("second: %w", fmt.Errorf("first: %w", errFmt.New()))
	)

	tests := []struct {
		name string
		e    error
		want error
	}{
		{"error chain", chain, err},
		{"error chain fmt", chainFmt, errFmt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err Comparer
			if got := errors.As(tt.e, &err); !got || err.Error() != tt.want.Error() {
				t.Errorf("errors.As() = %v; want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkWrap(b *testing.B) {
	const strerr = String("error %f")

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var err = errors.New("cause")
		err = strerr.New().Wrap(err)
		_ = err.Error()
		_ = errors.Is(err, strerr)
	}
}

func BenchmarkWrapWithStack(b *testing.B) {
	const strerr = String("error")

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var err = errors.New("cause")
		err = strerr.New().WithStack()
		_ = err.Error()
		_ = errors.Is(err, strerr)
	}
}

func BenchmarkErrorfWrap(b *testing.B) {
	var strerr = errors.New("error")

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var err = errors.New("cause")
		err = fmt.Errorf("%s: %w", strerr, err)
		_ = err.Error()
		_ = errors.Is(err, strerr)
	}
}
