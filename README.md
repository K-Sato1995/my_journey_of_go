# Printing

### General
`%v`:	the value in a default format
`%T`:	a Go-syntax representation of the type of the value

### Boolean
`%t`:	the word true or false

### Integer
`%b`:	base 2
`%o`:	base 8
`%d`:	base 10
`%x`:	base 16, with lower-case letters for a-f
`%X`:	base 16, with upper-case letters for A-F

### Floating-point
`%f`:	decimal point but no exponent,

### String
`%s`: the uninterpreted bytes of the string or slice
`%q`:	a double-quoted string safely escaped with Go syntax

### Pointer
`%p`: base 16 notation, with leading 0x

## References
- [Package fmt](https://golang.org/pkg/fmt/)
