package main

import "fmt"

// dclare a channel
var message = make(chan string)

func send() {
	// Send data using channel
	message <- "This is the message"
}

func hello(done chan bool) {
	fmt.Println("Hello GoRoutine")
	done <- true
}

func main() {

	// Sending and receiving data from a channel
	go send()
	msg := <-message
	fmt.Println(msg) //=> This is the message

	// Channesl lock
	ch := make(chan bool)

	go func() {
		fmt.Println("Hello")
		ch <- true
	}()

	fmt.Println(1)
	// It locks untile it gets a bool type data
	<-ch //=> Hello

	ch2 := make(chan bool)
	go hello(ch2)
	<-ch2 //=> Hello GoRoutine
}
