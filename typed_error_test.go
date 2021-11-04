package typed_errors

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
		formatted  = errFmt.WithArgs("formattedString")
		wrapped    = err.Wrap(errors.New("cause"))
		wrappedFmt = formatted.Wrap(errors.New("cause"))
		caused     = fmt.Errorf("caused by %w", err)
	)

	tests := []struct {
		name string
		e    TypedError
		args args
		want bool
	}{
		{"error has same type", err, args{err}, true},
		{"formattedString error", formatted, args{errFmt}, true},
		{"wrapped error", wrapped, args{err}, true},
		{"wrapped and formattedString", wrappedFmt, args{errFmt}, true},
		{"error was cause", wrapped, args{caused}, true},
		{"error has different type", err, args{String("other")}, false},
		{"error has different type but same text", err, args{errors.New("err")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Is(tt.args.err); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Wrap(t *testing.T) {
	type args struct {
		err   TypedError
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
		{"error wrap", args{err, cause}, "err: root"},
		{"error formattedString", args{errFmt.WithArgs("formattedString"), cause}, "err formattedString: root"},
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
		{"without cause", err.Wrap(nil), nil},
		{"with cause", err.Wrap(cause), cause},
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
		chain    = fmt.Errorf("second: %w", fmt.Errorf("first: %w", err))
		chainFmt = fmt.Errorf("second: %w", fmt.Errorf("first: %w", errFmt))
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
			var err TypedError
			if got := errors.As(tt.e, &err); !got || err.Error() != tt.want.Error() {
				t.Errorf("errors.As() = %v; want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkWrap(b *testing.B) {
	const strerr = ErrorString("error")

	b.ReportAllocs()
	var res string
	for i := 0; i < b.N; i++ {
		var err = errors.New("cause")
		err = strerr.Wrap(err)
		res = err.Error()
	}
	_ = res
}

func BenchmarkErrorfWrap(b *testing.B) {
	var (
		strerr = errors.New("error")
		res    string
	)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var err = errors.New("cause")
		err = fmt.Errorf("%s: %w", strerr, err)
		res = err.Error()
	}
	_ = res
}
