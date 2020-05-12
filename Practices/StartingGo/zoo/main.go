package main

import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
	fmt.Println(AppName())
}

// go run *.go
// =>
// Grass
// Banana
// Carrot
// Animal App
