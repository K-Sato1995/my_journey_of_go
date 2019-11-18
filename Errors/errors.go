package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")

	if err != nil {
		fmt.Println("An error has occured!")
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "Opened successfully!")
}
