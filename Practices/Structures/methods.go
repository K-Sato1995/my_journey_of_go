package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (h Person) intro(greeting string) string {
	return greeting + h.name
}

func NewPerson(name string, age int) *Person {
	person := new(Person)
	person.name = name
	person.age = age
	return person
}

func main() {
	ken := NewPerson("ken", 12)
	fmt.Println(ken.intro("Hello")) ///=> HelloKen
}
