package main

import "fmt"

func main() {
	var pointer1 *int
	var pointer2 *string

	var str = "Go"
	var num = 10
	var num2 = 20

	pointer1 = &num
	pointer2 = &str
	pointer3 := &num2

	fmt.Println(pointer1, pointer2) //=> 0xc420016078 0xc42000e1e0
	fmt.Println(pointer3)           //=> 0xc420016080
	fmt.Println(&str)               //=> 0xc42000e1e0

	fmt.Println(*pointer1) //=> 10 //read var1 through the pointer.
	*pointer1 = 20         //change the value of var1 through the pointer.
	fmt.Println(num)       //=> 20 //the new value of var1
}
