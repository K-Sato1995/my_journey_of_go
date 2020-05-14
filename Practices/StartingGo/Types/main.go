package main

import "fmt"

func main() {
	num := 3
	// Type Cast
	byte := byte(num)

	fmt.Println(byte)

	// Rune
	r := 'あ'
	fmt.Println(r) //=> 12354

	// String
	s := "Hello World"
	fmt.Printf("%v, %T\n", s, s) //=> Hello World, string

	rs := `
	Hello!
	World!!
	`
	fmt.Println(rs)

	// Array
	a := [5]int{1, 2, 4, 5, 6}
	a2 := [...]string{"w"}

	fmt.Println(a)  //=> [1 2 4 5 6]
	fmt.Println(a2) //=> [w]
}
