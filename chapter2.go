package main

import "fmt"

func main(){
  var num = 0
  var sum = 1
  lang := "Ruby"

  fmt.Println(if_statement(num)) //=> The value is 0
  fmt.Println(condition(lang)) //=> This is not Go

  for i := 0; i < 5; i++ {
      num += i
  }

  for ;sum < 100; {
    sum += sum
  }

  fmt.Println(num, sum) //=> 10 128
}

func if_statement(arg int) string{
  if arg == 0 {
    return "The value is 0"
  } else {
    return "The value is not 0"
  }
}

func condition(arg string) string {
  if v:= "Go"; arg == v {
    return "This is Go"
  } else {
    return "This is not Go"
  }
}
