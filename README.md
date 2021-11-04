# go-typed-errors

## Why this repo was created?
Main reason was to create additional methods for better error typing support.

## Features
- Errors as constants
- `errors.Is` support
- `xerrors.Wrapper` support with `Unwrap` method
- `String.NewWithArgs` support to add context arguments for error message, while `errors.Is` working
- `String.NewWithStack` support to store stack trace (untested)
- IDE highlighting, as type based on strings

### Show me the code

https://play.golang.org/p/ZSHBCxXQx6A

```go
package main

import (
	"errors"
	typedErrors "github.com/Bogdan-D/go-typed-errors"
	"log"
)

const ErrWrongBehaviour = typedErrors.String("wrong behaviour: %v")

func someFunc() error {
	return errors.New("someFunc failed")
}

func typedError(args ...interface{}) error {
	err := someFunc()
	if err != nil {
		return ErrWrongBehaviour.NewWithArgs(args...).Wrap(err)
	}
	return nil
}

func main() {
	err := typedError("not typed errors")

	// handle ErrWrongBehaviour error type
	if errors.Is(err, ErrWrongBehaviour) {
		log.Print("typedError handled: ", err)
	} else if err != nil {
		log.Print("other error cases:", err)
	}
}

```

### Benchmark
Comparsion `Wrap` with `Errorf` for wrapping errors

```
goos: linux
goarch: amd64
pkg: github.com/Bogdan-D/go-typed-errors
cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
BenchmarkWrap-8                  7856391               180.7 ns/op            88 B/op          5 allocs/op
BenchmarkErrorfWrap-8            4080972               291.3 ns/op            64 B/op          3 allocs/op
```
