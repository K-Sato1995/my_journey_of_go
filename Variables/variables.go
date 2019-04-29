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

	// Declaring mutiple variables at a time
	var (
		str3 = "JS"
		num5 = 4
		arr  = [2]string{"Go", "Ruby"}
	)
	fmt.Printf("%v, %v, %v\n", str3, num5, arr) //=> JS, 4, [Go Ruby]
	fmt.Printf("%T, %T, %T\n", str3, num5, arr) //=> string, int, [2]string
}
