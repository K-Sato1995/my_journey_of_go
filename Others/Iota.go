package main

import "fmt"

const (
	NUM1 = iota
	NUM2
	NUM3
	NUM4
	NUM5
)

func main() {
	fmt.Printf("%v,%v,%v,%v,%v", NUM1, NUM2, NUM3, NUM4, NUM5) //=>> 0,1,2,3,4
}
