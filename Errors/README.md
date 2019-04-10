# Errors

Go programs express error state with `error` values.
The `error` type is a built-in interface.

```go
type error interface {
  Error() string
}
```

Just like any other built in type such as `int`, `float64`..., `error` values can be stored in variables, returned from functions and so on.

The idiomatic way of handling error in Go is to compare the returned error to nil. A nil value indicates that no error has occurred and a non nil value indicates the presence of an error.

```go
err != nil
```

# References
- [GOLANGBOT.COM](https://golangbot.com/error-handling/)
