package main

import "fmt"

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go receiver(ch)
	i := 0

	for i < 1000 {
		ch <- i
		i++
	}
}
