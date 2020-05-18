package main

import "fmt"

// For without any condition
func loop1() {
	i := 0
	fmt.Println("=============Loop1==============")
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}
}

// For with a condition
func loop2() {
	i := 0
	fmt.Println("=============Loop2==============")
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// Classic for loop
func loop3() {
	fmt.Println("=============Loop3==============")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// Continue
func loop4() {
	fmt.Println("=============Loop4==============")
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
		i++
	}
}

// Range
func loop5() {
	items := [3]string{"Item1", "Item2"}
	fmt.Println("=============Loop5==============")
	for i, s := range items {
		fmt.Printf("%v[%v]\n", s, i)
	}
}
func main() {
	loop1()
	loop2()
	loop3()
	loop4()
	loop5()
}
