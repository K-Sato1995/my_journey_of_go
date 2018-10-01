package main

import "fmt"

func main() {
  arry := [6] int { 1, 2, 3, 4, 5, 6 }

  var slice1 []string
  var slice2 []string = []string { "Go", "Ruby" }
  slice3 := arry[0:2]
  slice4 := arry[0:]
  slice5 := arry[:4]
  slice6 := arry[:]
  slice7 := make([]string, 2, 2)

  fmt.Println(slice1) //=> []
  fmt.Println(slice2) //=> [Go Ruby]
  fmt.Println(slice3) //=> [1 2]
  fmt.Println(slice4) //=> [1 2 3 4 5 6]
  fmt.Println(slice5) //=> [1 2 3 4]
  fmt.Println(slice6) //=> [1 2 3 4 5 6]
  fmt.Println(slice7) //=> [ ]
}
