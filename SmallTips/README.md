# Go Small Tips

# Unreferenced package

You can keep unused packages below without causing any error by adding `_` to the packages.

```go
import _"fmt"
```

# Define variables or constants all together

You can define multiple variables or constants all together like the code below.

```go
var (
  n = 1
  s = "string"
)

const (
  t = true
  T = false
)
```

# Unnamed function

```go
f := func(x, y int) int { return x + y }
f(3,5) //=> 8
```

# Using an aliases for packages

You can make an aliases for packages like the code below.

```go
import (
  f "fmt"
  . "test" // You can omit the package name when you use anything from this package.
)

func main() {
  f.Println("hello")
  T // From test
}
```

# You must pass a bool value to if's condition expression

the value of condition expression in a `if statement` has to be a bool type.

```go
if(true) {} // ok

if(1){} // 条件式が論理値(bool型)ではないのでコンパイルエラーを引き起こす。
```

# iota

Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers.

```go
const (
  A = iota // A == 0
  B = iota // B == 1
  C = iota // C == 2
  )
```

You can start the number from 1 like the below.

```go
const (
  A = 1 + iota // A == 1
  B // B == 2
  C // C == 3
)
```

# Swapping values in a slice

You can swap values in a slice like the code below.

```go
func main() {
  var list = []int { 1, 2, 3 ,4,5}
  list[0], list[1] = list[1], list[0] // 逆の順に並べる。
  fmt.Println(list) //=> [2 1 3 4 5]
}
```
