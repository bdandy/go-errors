# go-typed-errors

## Why this repo was created?
Main reason was to create additional methods for better error typing support.

## Features
- Errors as constants
- `errors.Is` support
- `Wrap` method to wrap original error with `errors.Unwrap` method
- `String.NewWithArgs` support to add context arguments for error message, while `errors.Is` still compares error itself
- `String.NewWithStack` support to store stack trace (untested)

### Show me the code

https://play.golang.org/p/VS1UWYi5AY6

```go
package main

import (
	"errors"
	"log"

	typedErrors "github.com/Bogdan-D/go-typed-errors"
)

const ErrSomeFunc = typedErrors.String("somefunc for %s failed")

// let's image someFunc is 3rd-party dependency
func someFunc() error {
	return errors.New("io error")
}

func funcWithArgs(args ...interface{}) error {
	err := someFunc()
	if err != nil {
		return ErrSomeFunc.NewWithArgs(args...).Wrap(err)
	}
	return nil
}

func main() {
	err := funcWithArgs("tryme!")

	// handle ErrSomeFunc error type
	if errors.Is(err, ErrSomeFunc) {
		log.Print("typedError handled: ", err)
	} else if err != nil {
		log.Print("other error cases:", err)
	}
}
```

### Benchmark
Comparsion with `errors.Errorf` and `pkg/errors`

```
goos: linux
goarch: amd64
pkg: github.com/Bogdan-D/go-typed-errors
cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
BenchmarkWrap-8                          6708042               212.2 ns/op           160 B/op          5 allocs/op
BenchmarkWrapWithStack-8                 1345825               940.3 ns/op           421 B/op          7 allocs/op
BenchmarkErrorfWrap-8                    4218820               284.9 ns/op            64 B/op          3 allocs/op
BenchmarkPkgErrorWrap-8                  1376254               858.2 ns/op           368 B/op          6 allocs/op
BenchmarkPkgErrorWrapWithStack-8          781378              1593 ns/op             672 B/op          9 allocs/op
PASS
```
