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
variable := <- c // read from channel ch
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
