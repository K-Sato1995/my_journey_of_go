package main

import "fmt"

func main() {
	var sum int

	// Loop
	for i := 0; i < 5; i++ {
		sum += i
	}

	fmt.Printf("The sum is %d\n", sum) //=> The sum is 10

	// For loop used like the while loop in JS/C
	// Ommiting the init and post statements.
	number := 5
	for number < 100 {
		number += number
	}
	fmt.Printf("The number is %d\n", number) //=> The number is 160
}
