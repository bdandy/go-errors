# go-typed-errors

## Why this repo was created?
Main reason was to create additional methods for better error typing support.

## Features
- `errors.Is` support
- `xerrors.Wrapper` support with `Unwrap` method
- `#WithArgs` support to add context arguments for error message, while `errors.Is` still working
- Errors as constants
- IDE highlighting, as type based on strings
-  ... Ask for new features.


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
		return ErrWrongBehaviour.WithArgs(args...).Wrap(err)
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
