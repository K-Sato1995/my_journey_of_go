package main

import "fmt"

func main() {
	// Variable declaration
	var num1, num2 int
	fmt.Printf("%d, %d\n", num1, num2) //=>0, 0

	// Variables with initializers
	var str = "Ruby"
	var num3 = 2
	fmt.Printf("%s: %d\n", str, num3) //=> Ruby: 2

	// Short variable declarations
	str2 := "Go"
	num4 := 2
	fmt.Printf("%s: %d\n", str2, num4) //=> Go: 2
}
