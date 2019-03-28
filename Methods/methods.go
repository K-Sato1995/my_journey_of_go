package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

// Define a method
func (this_person Person) intro(greetings string) string {
	return "My name is " + this_person.name + " I'm " + strconv.Itoa(this_person.age) + " " + greetings
}

type Num struct {
	x int
	y int
}

func (p Num) stay() {
	p.x = 10
	p.y = 10
}

func (p *Num) modify() {
	p.x = 10
	p.y = 10
}

func main() {
	me := Person{"Mike", 24}
	fmt.Println(me.intro("Greetings everyone!")) //=> My name is Mike I'm 24 Greetings everyone!

	// Pointer receivers
	num1 := Num{0, 0}
	num1.stay()
	fmt.Println(num1) //=> {0 0}

	num1.modify()
	fmt.Println(num1) //=> {10 10}
}
