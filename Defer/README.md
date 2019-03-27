# Defer
 A `defer` statement defers the execution of a function until the surrounding function returns.
 For instance, `fmt.Println("Hello")` is excuted after `fmt.Println("World")` due to the `differ` statement that was given in the example below.

```go
package main

import "fmt"

func main(){
  defer fmt.Println("Hello")

  fmt.Println("World")
  //=> World
  //   Hello
}
```

 Deferred function calls are executed in `last-in-first-out` order. That means the first function that is given `defer` would be excuted lastly.

```go
package main

import "fmt"

func main(){
  defer fmt.Println(1)
  defer fmt.Println(2)
  defer fmt.Println(3)
  //=> 3
  //   2
  //   1
}
```
