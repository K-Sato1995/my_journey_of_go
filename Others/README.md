# Iota

Go's `iota` identifier is used in const declarations to simplify definitions of incrementing numbers. Because it can be used in expressions, it provides a generality beyond that of simple enumerations.

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

# References
- [Iota](https://github.com/golang/go/wiki/Iota)
