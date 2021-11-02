# go-typed-errors

## Why this repo was created?
Main reason was to create additional methods for better error typing support.

## Why not to use `errors.As` and custom structs? 
Who says I'm not using them? :)

### Compare (real go-sock4 changes):
```go
	var socksErr socks4.ErrorCause
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
	// skip ErrIdentRequired error type
	if err != nil && !errors.Is(err, socks4.ErrIdentRequired) {
		t.Error(err)
	}

```

