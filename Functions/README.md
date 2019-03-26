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
