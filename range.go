package main

import "fmt"

var arry_ex  = [3]string { "Go", "Ruby", "Python" }
var slice_ex = []string { "Go", "Ruby", "Python" }
var map_ex = map[string]string{ "Name":"Sam", "Gender":"Male" }

func main(){
  for index, value := range arry_ex {
    fmt.Println(index, value)
    //=> 0 Go
    //=> 1 Ruby
    //=> 2 Python
  }

  for index, ele := range slice_ex {
    fmt.Println(index, ele)
    //=> 0 Go
    //=> 1 Ruby
    //=> 2 Python
  }

  for key, value := range map_ex {
    fmt.Println(key, value)
    //=> Name Sam
    //=> Gender Male
  }

  for _, value := range arry_ex {
    fmt.Println(value)
    //=> Go
    //=> Ruby
    //=> Python
  }

  for index, _ := range arry_ex {
    fmt.Println(index)
    //=> 0
    //=> 1
    //=> 2
  }
}
