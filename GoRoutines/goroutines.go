package main

import (
	"fmt"
	"runtime"
	// "time"
	"log"
)

func hello() { fmt.Println("Hello World") }

func main() {
	go hello()
	// time.Sleep(1 * time.Second)
	fmt.Println("Main function")
	log.Println(runtime.NumGoroutine()) //=> 2019/03/31 14:11:47 2
	//=> main function
}
