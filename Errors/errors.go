package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")

	if err != nil {
		fmt.Println("The error is gonna occur!")
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "Opened successfully!")
}
