package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	fmt.Println("Start!")

	go func() {
		process(1, "A")
		ch1 <- true
	}()

	go func() {
		process(1, "B")
		ch2 <- true
	}()

	<-ch1
	<-ch2
	fmt.Println("Finish!")
}

func process(num int, str string) {
	for i := 0; i <= num; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, str)
	}
}
