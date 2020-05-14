package main

import "fmt"

// Normal function
func plus(x, y int) int {
	return x + y
}

// No Return value
func noReturnValue(word string) {
	fmt.Printf("%T", word)
	return
}

// Multiple Return Values
func multipleReturnValues(x, y string) (string, string) {
	return y, x
}

// Anonymous function
var minus = func(x, y int) int {
	return x - y
}

// Function that returns another function
func returnFunc() func() {
	return func() {
		fmt.Println("Returning another func")
	}
}

// Function that takes another function as an arg
func callFunction(function func()) {
	function()
}

// Closure
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	total := plus(3, 5)

	fmt.Println(total) //=> 8

	noReturnValue("Yo\n") //=> string

	fmt.Println(multipleReturnValues("Hello", "World")) //=> World Hello

	st1, _ := multipleReturnValues("Hello", "World")
	fmt.Println(st1) //=> World

	fmt.Println(minus(10, 4)) //=> 6

	returnFunc()() //=> Returning another func

	callFunction(func() {
		fmt.Println("I am a function")
	}) //=> I am a function

	fmt.Println(counter())
	fmt.Println(counter())
}
