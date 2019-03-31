package main

import "fmt"

var arr = [3]string {"Go", "Ruby", "Python"}
var slice = []string { "Go", "Ruby"}
var map1 = map[string]string { "Key1": "Value", "Key2": "Value"}

func main() {
  for index, value := range arr {
    fmt.Println(index, value)
  }//=> 0 Go 1 Ruby 2 Python

  for index, value := range slice {
    fmt.Println(value, index)
  } //=> Go 0 Ruby 1

  for key, value := range map1 {
    fmt.Println(key, value)
  } //=> Key1 Value Key2 Value

  // Skipping the index, key or value
  // Index only
  for index, _ := range arr {
    fmt.Println(index)
  }//=> 0 1 2

  // Value only
  for _, value := range slice {
    fmt.Println(value)
  } //=> Go Ruby
}
