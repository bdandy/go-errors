# go-errors

## Why this repo was created?
My own implementation of sentinel errors (https://dave.cheney.net/tag/errors)

## Features
- `Errors` are constants
- `errors.Is` support
- `Wrap` method to wrap original error with `errors.Unwrap` method
- `String.New` support to add context arguments for error message, while `errors.Is` still compares error itself
- `Error.WithStack` support to store stack trace at time method called


## Way to go with errors

My recommendation is to design your packages to return all errors which are declared. 
So you have to declare package level error, and wrap original error into your sentinel error.

For example, if your package works with filesystem, consumer don't know anything about `io.EOF` error.
`io.EOF` should be wrapped into your custom error type `EOF` and returned to consumer, so he won't search 3rd party packages errors.

### Show me the code

```go 
package mypkg

import (
    "io"
    "errors"
    
    serr "github.com/bdandy/go-errors"
)

const (
    ErrEOF = serr.String("EOF")
    ErrUnknown = serr.String("unknown error")
)

func testfunc() error {
    return io.EOF
}

// TestEOF returns ErrEOF or ErrUknown on some unexpected cases
//  io.EOF is wrapped into ErrEOF so mypkg has all expected errors in one place,
//  so it's more easy to use mypkg and handle errors 
//  (you don't have to know anything about `io` package and it's behaviour)
func TestEOF() error {
    err := testfunc() 
    
    switch {
    case errors.Is(io.EOF):
        return ErrEOF.New().Wrap(err)
    default: 
        return ErrUknown
    }
}

```

### Benchmark
Comparison with `errors.Errorf` and `pkg/errors`

```
goos: linux
goarch: amd64
pkg: github.com/bdandy/go-errors
cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
BenchmarkWrap-8                  	 5007164               270.9 ns/op           136 B/op          6 allocs/op
BenchmarkWrapWithStack-8         	 1232276               947.3 ns/op           392 B/op          7 allocs/op
BenchmarkErrorfWrap-8                    4218820               284.9 ns/op           64 B/op           3 allocs/op
BenchmarkPkgErrorWrap-8                  1376254               858.2 ns/op           368 B/op          6 allocs/op
BenchmarkPkgErrorWrapWithStack-8          781378               1593 ns/op            672 B/op          9 allocs/op
PASS
```
