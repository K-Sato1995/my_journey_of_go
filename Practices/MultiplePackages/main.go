package main

import (
	"./fruit"
	"./fruit/sub"
	"fmt"
)

func main() {
	fmt.Println("Hello from main.go")
	fruit.Hello()           //=> Hello From Fruit!
	fmt.Println(sub.Fruit1) //=> Apple
}
