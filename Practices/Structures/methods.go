package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (h Person) intro(greeting string) string {
	return greeting + h.name
}

func main() {
	ken := Person{"Ken", 23}

	fmt.Println(ken.intro("Hello")) ///=> HelloKen
}
