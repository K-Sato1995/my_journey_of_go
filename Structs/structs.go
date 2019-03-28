package main

import "fmt"

type person struct {
  name string
  age int
}

func main() {
  // Struct Instantiation
  // (1) The var keyword and Dot Notation
  var mike person
  mike.name = "Mike"
  mike.age = 23
  intro(mike.name, mike.age) //=> My name is Mike. I'm 23

  // (2) The var keyword and := operator
  bob := person { "Bob", 30 }
  sam := person { age: 35, name: "Sam" }
  intro(bob.name, bob.age) //=> My name is Bob. I'm 30
  intro(sam.name, sam.age) //=> My name is Sam. I'm 35

  // (3) Using the new keyword
  jen := new(person)
  jen.name = "Jennifer"
  jen.age = 23

  intro(jen.name, jen.age) //=> My name is Jennifer. I'm 23

  // Structs and Pointers
  var john person
  p_pointer := &john

  p_pointer.name = "John"
  p_pointer.age = 15

  fmt.Printf("%p\n", p_pointer) //=> 0xc42009a020
  intro(john.name, john.age) //=> My name is John. I'm 15
}

func intro(name string, age int) {
  fmt.Printf("My name is %s. I'm %d\n", name, age)
}
