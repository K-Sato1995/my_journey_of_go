# What are Goroutines?
 A `goroutine` is a lightweight thread managed by the Go runtime.

# How to start a Goroutine
 You can declare a `goroutine` by adding the keyword `go` before a function or method invocation like the code below.

```go
go function_name(arg, ...)
```
 In the following code, `go hello()` starts a new `goroutine` and `hello()` will run concurrently along with the `main` function which runs in its own Goroutine and its called the `main Goroutine`

```go
package main

import "fmt"

func hello() {
  fmt.Println("Hello World")
}

func main() {
  go hello()
  fmt.Println("main function")
  //=> main function
}
```

 The code above only outputs `main function`. This is due to the following reasons.  
 Firstly, when a new `goroutine` starts running, the goroutine call returns immediately. It does not wait for the gorutine to finish exuting.  
 Secondly, the `main Goroutine` should be running for any other `goroutines` to run. If the `main goroutine` terminates, the program will be terminated and no other `goroutine` will run.  
 Now you understand why it did'nt output `Hello World`. In the code, after calling `go hello()`, the control returned immediately to the next line of code which is `fmt.Println("main function")` without waiting for the `hello goroutine` to finish.

 You can excute `hello()` by giving it enough time to run before the `main goroutine` terminates.

```go
package main

import (
    "fmt"
    "time"
)

func hello(){
    fmt.Println("Hello World")
}

func main(){
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")

    //=> Hello World
    //=> main function
}
```
  Normaly, [channels]() are used to block the `main goroutine` until all other `goroutines` finish their execution.

# How to check the number of all existing goroutines
 You can use `runtimeNumGoroutine()` to get all the existing goroutines in your code.

```go
package main

import (
     "log"
     "runtime"
     "fmt"
    )

func hello(){
    fmt.Println("Hello World")
}

func main(){
    go hello()
    fmt.Println("main function")
    log.Println(runtime.NumGoroutine()) //-> 2018/10/09 07:39:07 2
}
```

# More Resources
[GOLANGBOT.COM](https://golangbot.com/goroutines/)
[Go Resources (Concurrency)](https://www.golang-book.com/books/intro/10)
[A Tour of Go (Goroutines)](https://tour.golang.org/concurrency/1)
