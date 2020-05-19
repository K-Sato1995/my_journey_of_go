package main

import "fmt"

func biggerThanTen(x int) string {
	if x > 10 {
		return fmt.Sprintf("%d is bigger than 10", x)
	} else {
		return fmt.Sprintf("%d is smaller than 10", x)
	}
}

func switchFunc() {
	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("Unknown")
	}
}

func switch2() {
	n := 3
	switch {
	case n > 4:
		fmt.Println("Bigger than 4")
	default:
		fmt.Println("Smaller than 4")
	}
}
func main() {
	fmt.Println(biggerThanTen(4)) //=> 4 is smaller than 10
	switchFunc()                  //=> 3 or 4
	switch2()                     //=> Smaller than 4
}
