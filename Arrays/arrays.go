package main

import "fmt"

func main() {
	// Define arrays

	// (1) var name[num]Type
	var langs [2]string
	langs[0] = "Go"
	langs[1] = "Ruby"
	fmt.Printf("%v\n", langs) //=> [Go Ruby]

	// (2) var name[num]Type = [num]Type { ele1, ele2, elen.... }
	var langs2 [2]string = [2]string{"Go", "Ruby"}
	fmt.Printf("%v\n", langs2) //=> [Go Ruby]

	// (3) name := [...] Type { ele1, ele2, elen... }
	var langs3 = [...]string{"Go", "Ruby"}
	fmt.Printf("%v\n", langs3) //=> [Go Ruby]

	// Accesing an array
	fmt.Printf("The first value is %v\n", langs[0]) //=> The first value is Go

	// Change an element of an array
	langs[0] = "Javascript"
	fmt.Printf("The first value is %v\n", langs[0]) //=> The first value is Javascript

	// Show the length of an array
	fmt.Printf("The length is %v\n", len(langs)) //=> The length is 2
}
