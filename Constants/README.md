# Constants
Constants are declared with the `const` keyword.  
Go supports constants of character, string, boolean, and numeric values.  
__Constants cannot be declared using the `:=` syntax__.

```go
package main

import "fmt"

const Name = "Go"

func main() {
	const Num = 1
	const Truth = true
	fmt.Println(Name, Num, Truth) //=> Go 1 true
}
```
