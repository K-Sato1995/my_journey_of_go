# Pointers

### How to declare a pointer
 In brief, a `pointer` is a value which points to the memory address of another value. A `pointer` can be declared using the `*` followed by the type of the stored value in Go.      
The variable `pointer1` is a pointer to an `int` type value and `pointer2` is a pointer to a `string` type value in the following code.

```go
package main

import "fmt"

func main() {
  var pointer1 *int
  var pointer2 *string

  fmt.Println(pointer1, pointer2) //=> nil nil // A pointer's zero value is `nil`
}
```

### How to retrieve the memory address of a variable
 In Go, you can use the `&` operator to retrieve the memory address of a variable. In the following code, `&str` represents the memory address of the variable `str`.

```go
package main

import "fmt"

func main() {
  var str = "Go"
  fmt.Println(&str) //=> 0xc42000e1e0
}
```

### How to use pointers
 You can assign the memory address of a variable to a `pointer` just like the code below.

```go
package main

import "fmt"

func main() {
  var pointer1 *int
  var pointer2 *string

  var str = "Go"
  var num = 10

  pointer1 = &num
  pointer2 = &str

  fmt.Println(pointer1, pointer2) //=> 0xc420016078 0xc42000e1e0
}
```

 You can also declare a `pointer` more succinctly like the following code.

```go
package main

import "fmt"

func main() {
  var num2 = 20
  pointer3 := &num2

  fmt.Println(pointer3) //=> 0xc420016080
}
```

 `pointer1` and `pointer3` can only store the memory address of `int` type values and `pointer2` can only store the memory address of a `string` type value respectively.

```go
package main

import "fmt"

func main() {
	var pointer1 *int
	var str = "Go"

	pointer1 = &str //=> cannot use &str (type *string) as type *int in assignment
}
```

### Accessing the value of a variable through its pointer
 The `*` operator denotes the pointer's underlying value. By using this feature, you can set, read and change the value of a variable through its pointer.

```go
package main

import "fmt"

func main() {
  var pointer1 *int
  var num = 10

  pointer1 = &num

  fmt.Println(*pointer1) //=> 10 //read var1 through the pointer.
  *pointer1 = 20 //change the value of var1 through the pointer.
  fmt.Println(num) //=> 20 //the new value of num
}
```
