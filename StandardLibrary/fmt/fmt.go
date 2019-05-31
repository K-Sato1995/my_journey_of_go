package main

import "fmt"

func main() {
	// Print + new line
	fmt.Println("Hello World")
	// => Hello World

	fmt.Printf("decimal=%d, binary=%b, octal=%o, hexadecimal=%x\n", 17, 17, 17, 17)
	// => decimal=17, binary=10001, octal=21, hexadecimal=11

	// %v %#v %T can be very useful for debugging
	ar := [...]int{1, 2, 3}

	// Show each value
	fmt.Printf("Digit=%v, String=%v, Array=%v\n", 1, "String", ar)
	// => Digit=1, String=String, Array=[1 2 3]

	// Show the literal expression of each value
	fmt.Printf("Digit=%#v, String=%#v, Array=%#v\n", 1, "String", ar)
	// => Digit=1, String="String", Array=[3]int{1, 2, 3}

	// Show the type of each value
	fmt.Printf("Digit=%T, String=%T, Array=%T\n", 1, "String", ar)
	// => Digit=int, String=string, Array=[3]int
}
