package main

import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan bool)

  fmt.Println("Start!")
  go process(1, "A", ch)
  <-ch
  go process(1, "B", ch)
  <-ch
  fmt.Println("Finish!")
}


func process(num int, str string, ch chan bool) {
  for i:= 0; i <= num; i++ {
    time.Sleep(1 * time.Second)
    fmt.Println(i, str)
  }
  ch <- true
}
