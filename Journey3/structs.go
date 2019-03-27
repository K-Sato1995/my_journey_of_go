package main

import "fmt"

type person struct {
	first_name string
	age        int
}

func main() {

	//The `var` keyword and Dot Notation
	var mike person
	// mike.first_name = "Mike"
	// mike.age = 20

	//The `var` keyword and `:=` operator
	bob := person{"Bob", 30}
	sam := person{age: 40, first_name: "Sam"}

	//Using the `new` keyword
	jen := new(person)
	jen.first_name = "Jennifer"
	jen.age = 10

	s_pointer := &mike
	fmt.Println(s_pointer)
	s_pointer.first_name = "Mike"
	s_pointer.age = 20
	fmt.Println(mike) //=> {Mike 20}

	fmt.Println(jen.first_name, jen.age)   //=> Jennifer 10
	fmt.Println(sam.first_name, sam.age)   //=> Sam 40
	fmt.Println(bob.first_name, bob.age)   //=> Bob 30
	fmt.Println(mike.first_name, mike.age) //=> Mike 20
}
