package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch

		if ok == false {
			// 受信できなくなったら終了
			break
		}
		fmt.Println(name, i)
	}
	fmt.Printf("%s is done\n", name)
}

func main() {
	ch := make(chan int, 20)

	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}

	time.Sleep(3 * time.Second)
}
