package main

import "fmt"

func main(){
  // Declare a pointer
  var pointer1 *int
  var pointer2 *string

  // Default value of a pointer is nil
  fmt.Println(pointer1, pointer2) //=> <nil> <nil>

  // Retrieve the memory address of a variable
  var str = "Go"
  var num = 1
  fmt.Println(&str) //=> 0xc42000e1e0
  fmt.Printf("The memory address is %p\n",&str) //=> The memory address is 0xc42000e1e0

  // Assign the memory address of a variable to a pointer
  pointer1 = &num
  pointer2 = &str

  // Accessing the value of a variable through its pointer
  fmt.Printf("The number is %d", *pointer1)

  //change the value of a value through the pointer.
  *pointer2 = "Ruby"
  fmt.Printf("The string is %s\n", *pointer2) //=> The number is 1The string is Ruby
}
