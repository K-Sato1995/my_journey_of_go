package main

import "fmt"

func main() {
	arry := [...]string{"Go", "Ruby", "Javascript"}

	slice := arry[:1]
	var new_slice = append(slice, "Java", "Swift", "C")

	var nil_slice []int
	if nil_slice == nil {
		fmt.Println("nil") //=> nil
	}

	fmt.Println(new_slice) //=> [[Go Ruby Javascript Java Swift C]

	fmt.Println(len(slice), cap(slice)) //=> 1 3
	fmt.Println(slice)                  //=> [Go Ruby Javacript]

	slice[0] = "Python"

	fmt.Println(slice) //=> [Python Ruby Javascript]
	fmt.Println(arry)  //=> [Python Ruby Javascript]
}
