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
