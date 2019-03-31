package main

import (
	"fmt"
	"math"
	"math/rand"
)

var var1, var2, var3 bool
var str1 = "ruby"
var str2, num1, num2 = "go", 2, 4

const Name = "Go"

func main() {
	var num int
	str3 := "Js"
	num3 := 4
	const Num = 1
	const Truth = true

	fmt.Println("My journey of Go", rand.Intn(10)) //=>My journey of Go 1
	fmt.Println(math.Pi)                           //=> 3.141592653589793
	fmt.Println(greetings("John"))                 //=> Hello John
	fmt.Println(add(2, 4))                         //=> 6
	a, b := multipleArgs("Hello", "World")
	fmt.Println(a, b)                   //=> World Hello
	fmt.Println(num, var1, var2, var3)  //=> 0 false false false
	fmt.Println(str1, num1, str2, num2) //=> ruby 2 go 4
	fmt.Println(num3, str3)             //=> 4 Js
	fmt.Println(Name, Num, Truth)       //=> Go 1 true
}

func greetings(name string) string {
	return "Hello" + " " + name
}

func add(x, y int) int {
	return x + y
}

func multipleArgs(arg1, arg2 string) (string, string) {
	return arg2, arg1
}
