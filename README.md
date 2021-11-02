# go-typed-errors

## Why this repo was created?
Main reason was to create additional methods for better error typing support.

## Why not to use `errors.As` and custom structs? 
Who says I'm not using them? :)

## Features
- `xerrors.Wrapper` support
- `errors.Is` support
- `errors.As` abandoned
- `#WithArgs` support to add context arguments for error message, while `errors.Is` still working
- Errors as constants
- Cool IDE highlighting, as type based on strings
-  ... Ask for new features.


### Compare
(from `go-socks4` source code)

```go
	// complex types definitions somewhere in other place...
	var socksErr socks4.Error
	if err != nil && errors.As(err, &socksErr) {
		switch {
		case socksErr.Equal(socks4.ErrIdentRequired):
		default:
			t.Error(err)
		}
	} else if err != nil {
		t.Error(err)
	}
```

### With
```go
	const (
		ErrIdentRequired = typedErrors.String("socks4 server require valid identd: %v")
	)

	// skip ErrIdentRequired error type
	if err != nil && !errors.Is(err, socks4.ErrIdentRequired) {
		t.Error(err)
	}

```

