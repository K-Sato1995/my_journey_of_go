package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("World")

	// Deferred function calls are executed in last-in-first-out order.
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// => 3 2 1 World Hello
}
