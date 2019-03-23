# Overview
 In this post, I will write about how looping and conditional statements work in Go.

# For
 `For` is Go's only looping construct. Basic `for loop` consists of three parts that are listed below.

- `init statement`: It is executed before the first iteration. (It is often used to define a variable that can only be used in the scope of the `for` statement.)

- `condition`: It is evaluated before every iteration. (If it's true, the loop body executes. otherwise, the loop terminates.)

- `post statement`: It is executed at the end of every iteration.

 The basic syntax of Go's `for loop` looks like this.

```go
  for init statement; condition; post statement {
    //content of the loop
  }
```

The loop terminates the iteration once the condition evaluates to `false`.

```go
package main

import "fmt"

func main(){
  var num = 0

  for i := 0; i < 5; i++ {
      num += i
  }

  fmt.Println(num) //=> 10
}
```

`init` and `post` statements can be omitted. When those statements are omitted, Go's `for loop` behaves like the `while loop` of Javascript/Java/C.

```go
package main

import "fmt"

func main() {
	num := 1
	for ;num < 100; {
		num += num
	}
	fmt.Println(num) //=> 128
}
```

 If you omit the `condition`, it loops infinitely.

```go
package main

func main() {
	for {
	}
}
```

# If
  The basic syntax of Go's `if` statement is similar to its `for` statement.

```go
package main

import "fmt"

func main(){
  var num = 0

  fmt.Println(if_statement(num)) //=> The value is 0
}

func if_statement(arg int) string{
  if arg == 0 {
    return "The value is 0"
  } else {
    return "The value is not 0"
  }
}
```

 Just like the `for` statement, `if` statement can start with a short statement that is executed before the condition.

```go
package main

import "fmt"

func condition(arg string) string {
  if v:= "Go"; arg == v {
    return "This is Go"
  } else {
    return "This is not Go"
  }
}

func main() {
	lang := "Ruby"
	fmt.Println(condition(lang)) ///=> This is not Go
}
```

# Switch
A `switch` statement is a shorter way to write a sequence of `if - else` statements. A `switch` statement only runs the first case that meets the condition, not all the cases that follow.

```go
package main

import "fmt"

func main(){
  lang := "Go"

  switch lang {
	case "Ruby":
	  	fmt.Println("This is Ruby")
	case "Go":
	  	fmt.Println("This is Go")
	default:
	  	fmt.Println("This is something else")
	}
	//=> This is Go
}
```

 You can also use the `init statement` with a `switch` statement.

```go
package main

import "fmt"

func main(){

  switch lang:= "Go"; lang {
	case "Ruby":
	  	fmt.Println("This is Ruby")
	case "Go":
	  	fmt.Println("This is Go")
	default:
	  	fmt.Println("This is something else")
	}
	//=> This is Go
}
```

`switch` without a condition is the same as `switch true`.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

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
