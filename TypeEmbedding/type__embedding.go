package main

import "fmt"

type Person struct {
  name string
}

type User struct {
  Person // Type Embedding
}

func(this_person Person) display_name() string  { //Person typed method
  return this_person.name
}

func main() {
 var bob = Person{ "Bob" }
 var mike = new(User)
 mike.name = "Mike"

 fmt.Println(bob.display_name()) //=> Bob
 fmt.Println(mike.display_name()) //=> Mike
}
