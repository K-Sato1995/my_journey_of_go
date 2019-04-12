# Iota

Go's `iota` identifier is used in const declarations to simplify definitions of incrementing numbers.Because it can be used in expressions, it provides a generality beyond that of simple enumerations.

```go
const (
	NUM1 = iota
	NUM2
	NUM3
	NUM4
	NUM5
)

func main() {
	fmt.Printf("%v,%v,%v,%v,%v", NUM1, NUM2, NUM3, NUM4, NUM5) //=>> 0,1,2,3,4
}
```

# Type Definitions

In Go, we can define new types by using the following form. In the syntax, type is a keyword.

```go
// Define a solo new type.
type NewTypeName SourceType

// Define multiple new types together.
type (
	NewTypeName1 SourceType1
	NewTypeName2 SourceType2
)
```

# References

- [Iota](https://github.com/golang/go/wiki/Iota)
- [Types](https://go101.org/article/type-system-overview.html)
