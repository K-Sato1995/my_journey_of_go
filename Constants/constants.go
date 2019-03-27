package main

import "fmt"

// Define a constant1
const Name string = "Jack"

func main() {
  // Define a constant2
  const age = 24

  // Use constants
  fmt.Printf("%s is %d years old.\n", Name, age) //=> Jack is 24 years old.
}
