package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var ken Person
	ken.name = "Ken"
	ken.age = 15
	fmt.Printf("I'm %s. I'm %d\n", ken.name, ken.age)

	jack := Person{"Jack", 34}
	fmt.Printf("I'm %s. I'm %d\n", jack.name, jack.age)

	bob := Person{name: "Bob", age: 3}
	fmt.Printf("I'm %s. I'm %d\n", bob.name, bob.age)
}
