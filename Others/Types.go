package main

import "fmt"

type NewType int

// Type Alias Declarations
type NewNewType = NewType

func main() {
	var num NewType
	var num2 NewNewType
	num = 2
	num2 = 5

	fmt.Printf("%v\n%d\n", num, num2) //=> 2 5
}
