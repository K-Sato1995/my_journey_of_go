package main

import "fmt"

func main() {
	var p *int
	num := 5
	p = &num
	fmt.Printf("p=%v num=%v\n", p, num) //=> p=0xc420016078 num=5
	fmt.Printf("*p = %v\n", *p)

	*p = 10
	fmt.Printf("p=%v num=%v\n", p, num) //=> p=0xc420016078 num=10
}
