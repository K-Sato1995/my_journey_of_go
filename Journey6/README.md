# Introduction
 In this post, I'll explain how `interfaces` work in Go.

# Interfaces in Go
 In Go, an `interface` is a set of method signatures. A variavle of interface type can hold any value that implements those methods.

# Basics
 You can declare an `interface` like the following code. It is basically just a list of methods.

```go
type NAME_OF_INTERFACE interface {
    method_name1(return_type1)    
    method_name2(return_type2)
    ..........
    method_namex(return_typex)
}
```
 Any type that provides the methods named in an interface is treated as an implementation of that interface. No explicit declaration is required.

# The empty interface
 The interface type that does not specify any method is known as the `empty interface`. It can be defined as the following code.

```go
interface {}
```
 A variable of empty interface type can hold values of any type.

```go
package main

import "fmt"

func main(){
  var intface interface{}
  intface = 1
  fmt.Println(intface) //=> 1

  intface = "string" //=> string
  fmt.Println(intface)

  intface = []string{"Go", "Ruby", "JS"}
  fmt.Println(intface) //=>[Go Ruby JS]
}
```

# Interface values
 A `interface value` is represented as a pair of a __concrete value__ and a __dynamic type__.

```go
[Value, Type]
```
You can use `%v` to print the concrete value and `%T` to print the dynamic type.

```go
package main

import "fmt"

func main(){
  var intface interface{}
  intface = 1
  fmt.Printf("%v %T\n", intface, intface) //=>1 int

  intface = "string"
  fmt.Printf("%v %T\n", intface, intface) //=> string string

  intface = []string{"Go", "Ruby", "JS"}
  fmt.Printf("%v %T\n", intface, intface) //=> [Go Ruby JS] []string
}
```

# Type assertions
A `type assertion` provides access to an interface value's underlying concrete value.

```go
concrete_value := Interface_value.(TYPE)
```
 This statement above asserts that `Interface_value` holds the concrete type `TYPE` and assigns the underlying `TYPE` value to `variable`.  
 To check whether an interface value holds a specific type, a `type assertion` can return two values. the `underlying value` and a `boolean value` that reports whether the assertion succeeded.

```go
concrete_value, ok := Interface_value.(TYPE)
```

```go
package main

import "fmt"

func main(){
  var intface interface{} = "Hello World"

  t := intface.(string)
  fmt.Println(t) //=> Hello World

  t2, ok := intface.(string)
  fmt.Println(t2, ok) //=> Hello World true

  t3, ok := intface.(float64)
  fmt.Println(t3, ok) //=> 0 false
}
```

# Type switches
 A `type switch` is a construct that permits several type assertions in series.  
 You can declare a `type switch` like the following code.

```go
switch v := x.(type) {
    case TYPE1:
	//here v has TYPE1
    case TYPE2:
	//here: v has TYPE2
    ...
    default: ...
}
```

 In the follwing code, the switch statement tests whether the interface value `i` holds a value of type `int` or `string`. In each of the `int` and `string` cases, the variable `v` will be of type `int` or `string` respectively and hold the value held by `i`.  
 In the default case (where there is no match), the variable `v` is of the same interface type and value as `i`.

```go
package main

import "fmt"

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	typeSwitch(21) //=> Twice 21 is 42
	typeSwitch("Hello World") //=> "Hello World" is 11 bytes long
	typeSwitch(true)//=> I don't know about type bool!
}
```

# Implementing interfaces
 `Interfaces` can be implemented as methods on structs like the following code.

```go
package main

import "fmt"

 type People interface {
  intro()
 }

 type Person struct {
   name string
 }

func (rarg Person) intro() string{
  return "Hello" + " " + "I'm" + " " + rarg.name
}

func main() {
  bob := Person{"Bob"}
  fmt.Println(bob.intro()) //=> Hello I'm Bob
}
```
