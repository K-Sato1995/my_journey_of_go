package main

import "fmt"

func main() {
	var arr [3]string
	arr[0] = "Go"
	arr[1] = "Ruby"
	arr[2] = "Javascript"

	var arr2 [3]string = [3]string{"Go", "Ruby", "Javascript"}

	arr3 := [...]string{"Go", "Ruby", "Javascript"}
	fmt.Println(arr, arr2, arr3) //=> [Go Ruby Javascript]
}
