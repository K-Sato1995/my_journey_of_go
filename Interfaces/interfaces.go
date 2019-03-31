package main

import "fmt"

var empty_interface interface {}

func main() {
  empty_interface = 1
  fmt.Printf("The type is %T\n", empty_interface) //=> The type is int
  empty_interface = "string"
  fmt.Printf("The type is %T\n", empty_interface) //=> The type is string
  empty_interface = []string{"Go"}
  fmt.Printf("The type is %T\n", empty_interface) //=> The type is []string
}
