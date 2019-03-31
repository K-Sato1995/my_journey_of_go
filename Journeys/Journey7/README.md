# Introduction
 This post is about `goroutines` and `channels`. This is the last post of  `My Journey of Go` and might have been the hardest one to write about since I was not familiar with the concepts at all.

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

# Channels
 `Channels` provide a way for `goroutines` to communicate with one another and synchronize their execution. Data can be sent from one goroutine and received from another goroutine.


# How to declare a channel
 You can declare a `channel` like the code below.

```go
ch := make(chan TYPE)
```
 Each `channel` has a type associated with it. This type is the type of data that the channel is allowed to transport. You can not transport any other data using the `channel`.

# Sending and receiving from a channel
 You can send and receive data from a `channel` using the syntax below.

```go
ch <- data // write to channel ch
variable := <- cd // read from channel ch
```
 In the first line, the arrow points towards `ch` and this means we are writing data to channel `ch`.  
 In the second line, the arrow points outwards from `ch` and this means we are reading from `ch` and storing data to `variable`.

```go
package main

import "fmt"

func main() {
  //Create a new channel
  ch := make(chan string)

  //Write data to ch
  go func() { ch <- "data" }()

  //Reading data from ch
  variable := <-ch
  fmt.Println(variable) //=> data
}
```

# Sends and receives are blocking by default
 By default, sends and receives block until the other side is ready. That means when a data is sent to a channel, the control is blocked until some other `goroutine` reads from that channel.  
 This allows `goroutines` to synchronize without explicit locks or condition variables.

### Example1
 In the following code, I created a boolean type channel named `ch` and declare a function as a goroutine. Then, I passed a boolean value to `ch` in the function.  
The `<-ch` statement will block the code until some boolean data is received on the `ch` channel.

```go
package main

import "fmt"

func main() {
  ch := make(chan bool)  //Create a boolean type channel

  //Run the following function as a goroutine.
  go func() {
    fmt.Println("Hello")

    ch <- true  //Pass a boolean value to the channel
  }()

  <- ch //Block until ch receives a boolean value
}
```

### Example2
 In this example, I defined a function called `hello` which accepts a `channel` as its argument. And then, I created a channel called `ch` in the `main goroutine` and passed it as a parameter to the `hello()` goroutine.  
 Like [Example1](# example1), the `<-ch` statement will block the code until some boolean data is received on the `ch` channel.

```go
package main

import "fmt"

func hello(channel chan bool) {  
    fmt.Println("Hello world")
    channel <- true
}
func main() {  
    ch := make(chan bool) //Create a boolean type channel

    go hello(ch) //Run hello method as a goroutine

    <-ch //Block until ch receives a boolean value
    fmt.Println("main function")

    //=> Hello world
    //   main function
}
```
# More Resources
[GOLANGBOT.COM](https://golangbot.com/goroutines/)
[Go Resources (Concurrency)](https://www.golang-book.com/books/intro/10)
[A Tour of Go (Goroutines)](https://tour.golang.org/concurrency/1)
