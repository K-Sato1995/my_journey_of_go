# Basic Types

## Rune type

A rune is a type meant to represent a Unicode code point.(It is actually just an alias for the type `int32`.)
It can be written as the following code.

```go
func main() {
	r := 'R'
	fmt.Printf("%d,  %T\n", r, r) //=> 82,  int32
}
```

# References

- [Go What is a rune?](https://programming.guide/go/rune.html)
