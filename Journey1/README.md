# What is 'My Journey of Go' for?
 `My Journey of Go` is a collection of posts about what I learned from [A Tour of Go](https://tour.golang.org/list). A Tour of Go is a fantastic place for all people to start their journey of Go programming language.

# Packages
 Every Go program consists of `packages`. Each package needs to be imported first to use its exported identifiers.  
 You can import `packages` like the code below.

```go
package main

import (
  "fmt"
  "math/rand"
)

func main(){
  fmt.Println("My journey of Go", rand.Intn(10)) //=>My journey of Go 1
}

```

  You can also import each package individually.

```go
import "fmt"
import "math/rand"
```

# Exported names
 In Go, a name is exported if it begins with a capital letter. For instance,  `Pi` is a name that is exported from `'math'` package.

```go
package main

import "math"

func main(){
  fmt.Println(math.Pi) //=> 3.141592653589793
}
```

# Functions
 A function can take zero or more arguments. You can define a function like the code below.

```go
func name_of_the_function(arguments) type {
  [content of the function]
}
```

 Don't forget to write the types of arguments and the return value of a function.

```go
package main

import "fmt"

func main(){
  fmt.Println(greetings("John")) //=> Hello John
}

func greetings(name string) string {
  return "Hello" + " " + name
}
```
### Omission of the type of arguments
When there are two or more parameters sharing the same type, you only have to write the type once.

```go
package main

import "fmt"

func main(){
  fmt.Println(add(2, 4)) //=> 6
}

func add(x, y int) int {
  return x + y
}
```

### Multiple results
A function can return any number of results.

```go
package main

import "fmt"

func main(){
  a, b := multipleArgs("Hello", "World")
  fmt.Println(a, b) //=> World Hello
}

func multipleArgs(arg1, arg2 string)(string, string) {
  return arg2, arg1
}
```

# Variables
 The `var` statement declares a list of variables. As in function argument lists, the type of a variable is written after the name of the variable.    
Variables that are declared without an initial value are given their zero value. For instance, `0` is for numeric types, `false` is for the boolean type and `""`(the empty string) is for strings.

```go
package main

import "fmt"

var var1, var2, var3 bool

func main() {
	var num int
	fmt.Println(num, var1, var2, var3) //=> 0 false false false
}
```
As you can see in the code above, A `var` statement can be used at package or function level.

### Variables with initializers
 A `var` declaration can include an initializer.  
 If there is an initializer, the type of the variable can be omitted. The variable will take the type of the initializer.

```go
package main

import "fmt"

func main() {
	var str1 = "ruby"
	var str2, num1, num2 = "go", 2, 4
	fmt.Println(str1,num1, str2, num2) //=>ruby 2 go 4
}
```

## Short variable declarations
 You can use `:=` to declare a variable inside a function. Go raises an error if you try to declare a variable with `:=` outside a function.

```go
package main

import "fmt"

func main() {
	str := "Js"
	num := 4
	boole := true

	fmt.Println(str, num, boole) //=> Js 4 true
}
```

### Basic types
 Go's basic types are listed in the code below.

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

### Type conversions
 You can not implicitly change the type of a variable in Go. Go raises an error if you try to reassign a value that type is different form the original value type like the code below.

```go
var i int = 100
var f float64 = i //=> // cannot use i (type int) as type float64
```

 If you want to change the type of a variable, you can do so like the code below.

```go
var a uint32 = 1234567890
var b uint8 = uint8(a)
fmt.Println(b)  // 210
```

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
