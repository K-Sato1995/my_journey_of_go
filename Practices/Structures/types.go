package main

import "fmt"

// Type
type AreaMap map[string][2]float64

// Type for functions
type CallBack func(i int) int

func main() {
	// Type
	amap := AreaMap{"Tokyo": {35.6488, 44.5}}
	fmt.Println(amap) //=> map[Tokyo:[35.6488 44.5]]

	// Type for functions
	ints := []int{2, 3}

	fmt.Println(Sum(ints,
		func(i int) int {
			return i * 2
		}))
}

func Sum(nums []int, callBack CallBack) int {
	var sum int
	for _, i := range nums {
		sum += i
	}
	return callBack(sum) //=> 10
}
