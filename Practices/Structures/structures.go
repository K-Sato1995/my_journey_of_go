package main

import "fmt"

// Address containes address info
type Address struct {
	coutry  string
	city    string
	zipCode int
}

// Person has name age and address
type Person struct {
	name    string
	age     int
	address Address
}

func main() {
	var ken Person
	address1 := Address{"USA", "LA", 233}
	ken.name = "Ken"
	ken.age = 15

	fmt.Printf("I'm %s. I'm %d\n", ken.name, ken.age, address1)

	jack := Person{"Jack", 34, address1}
	fmt.Printf("I'm %s. I'm %d\n", jack.name, jack.age)

	bob := Person{name: "Bob", age: 3, address: Address{"UAE", "TA", 2}}
	fmt.Printf("I'm %s. I'm %d\n", bob.name, bob.age)
}
