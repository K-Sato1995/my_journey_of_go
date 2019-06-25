package main

import (
	"fmt"
	"os"
)

func main() {
	result, err := os.Open("/nonexistence.txt")

	if err != nil {
		fmt.Println("There was an error")
		fmt.Println(err)
		return
		//=> There was an error
		//=> open /nonexistence.txt: no such file or directory
	}

	fmt.Println(result)
}
