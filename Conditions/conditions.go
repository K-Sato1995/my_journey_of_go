package main

import (
  "fmt"
  "time"
)

func main(){
  fmt.Println(ifStatement(2)) //=> The value is 2
  fmt.Println(ifWithInit("Go")) //=> true
  fmt.Println(switchStatement()) //=> 5
  swithWithInit() //=> Go
  switchWithoutCondition() //=> Hello
}


func ifStatement(arg int) string {
  // normal if statement
  if arg  == 0 {
    return "The valud is 0"
  }else if arg == 2 {
    return "The value is 2"
  }else {
    return "The value is not 0 or 2"
  }
}

func ifWithInit(arg string) bool {
  // If with the init statement.
  if lang := "Go"; arg == lang {
    return true
  } else {
    return false
  }
}


func switchStatement() string {
  //It only runs the first case that meets the condition
  var number = 5

  switch number {
    case 3:
      return "3"
    case 2:
      return "2"
    default:
      return "5"
  }
}

func swithWithInit() {
  // Swithc with the init statement.
  switch lang := "Go"; lang {
    case "Ruby":
      fmt.Println("Ruby")
    case "Go":
      fmt.Println("Go")
    default:
      fmt.Println("Something else")
  }
}

func switchWithoutCondition() {
  t := time.Now()

  // switch without a condition is the same as switch true
  switch {
  case t.Hour() < 12:
    fmt.Println("Good Morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon!")
  default:
    fmt.Println("Hello!")
  }
}
