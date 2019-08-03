# Printing

### General

- `%v`: the value in a default format
- `%T`: a Go-syntax representation of the type of the value

### Boolean

- `%t`: the word true or false

### Integer

- `%b`: base 2
- `%o`: base 8
- `%d`: base 10
- `%x`: base 16, with lower-case letters for a-f
- `%X`: base 16, with upper-case letters for A-F

### Floating-point

- `%f`: decimal point but no exponent,

### String

- `%s`: the uninterpreted bytes of the string or slice
- `%q`: a double-quoted string safely escaped with Go syntax

### Pointer

- `%p`: base 16 notation, with leading 0x

# Escape Sequence

- `\\`: \ character
- `\'`: ' character
- `\"`: " character
- `\?`: ? character
- `\a`: Alert or bell
- `\b`: Backspace
- `\f`: Form feed
- `\n`: Newline
- `\r`: Carriage return
- `\t`: Horizontal tab
- `\v`: Vertical tab

# Basic Types

- `bool`: false or true
- `string`: string

Numeric types:

- `uint`: either 32 or 64 bits
- `int`: same size as uint
- `uintptr`: an unsigned integer large enough to store the uninterpreted bits of a pointer value
- `uint8`: the set of all unsigned 8-bit integers (0 to 255)
- `uint16`: the set of all unsigned 16-bit integers (0 to 65535)
- `uint32`: the set of all unsigned 32-bit integers (0 to 4294967295)
- `uint64`: the set of all unsigned 64-bit integers (0 to 18446744073709551615)

- `int8`: the set of all signed 8-bit integers (-128 to 127)
- `int16`: the set of all signed 16-bit integers (-32768 to 32767)
- `int32`: the set of all signed 32-bit integers (-2147483648 to 2147483647)
- `int64`: the set of all signed 64-bit integers(-9223372036854775808 to 9223372036854775807)

- `float32`: the set of all IEEE-754 32-bit floating-point numbers
- `float64`: the set of all IEEE-754 64-bit floating-point numbers

- `complex64`: the set of all complex numbers with float32 real and imaginary parts
- `complex128`: the set of all complex numbers with float64 real and imaginary parts

- `byte`: alias for uint8
- `rune`: alias for int32 (represents a Unicode code point)

# Format your code

To format your code, you can use the gofmt tool directly:

```
gofmt -w yourcode.go
```

Or you can use the "go fmt" command:

```
go fmt path/to/your/package
```

# Naming Conventions

Using `MixedCaps` is the convention in Go.

```go
// var example
var testVariable string

// function example
testFunction (){}
```

But if you want to export a function, you have to write it starting with a capital letter like the name below.

```go
TestFunction () {}
```

# String interpolation

The Go `Sprintf` method from `fmt` package serves just this:

```go
import ( "fmt" )
date := fmt.Sprintf("%d-%d-%d", year, month, day)
time := fmt.Sprintf("%d:%d:%d", hour, minute, second)
datetime := fmt.Sprintf("%s,%s", date, time)
```

# References

- [Package fmt](https://golang.org/pkg/fmt/)
- [Go Data Types](https://www.tutorialspoint.com/go/go_data_types.htm)
- [go fmt your code](https://blog.golang.org/go-fmt-your-code)
- [Escape Sequence](https://www.tutorialspoint.com/go/go_constants.htm)
- [Awesome Go](https://github.com/avelino/awesome-go)
- [Naming](https://golang.org/doc/effective_go.html#names)
