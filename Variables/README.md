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

## Type conversion

The expression `T(v)` converts the value v to the type `T`.

```

```
