package main

import "fmt"

func main() {
	var list = []int{3, 5, 1, 2, 8, 6, 4}

	fmt.Println(quickSort(list))

}

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	// left is the value on the far left side.
	// right is the value on the far right side.
	left, right := 0, len(list)-1

	// Pick a pivot number (It can be any value in the list.)
	// rand.Int() % len(list) is used in the example.
	pivot := right

	list[pivot], list[right] = list[right], list[pivot]

	for i := range list {
		// if list[i] is smaller that list[right]
		// swap list[left] and list[i] and add 1 to left.
		if list[i] < list[right] {
			list[left], list[i] = list[i], list[left]
			left++
		}
		// Basically putting the smallest value to the far left side in every iteration till it reaches the last value in the list.
	}

	// fmt.Println(list)
	//=> [3 1 2 5 8 6 4]
	//=> [1 3 2]
	//=> [6 8]
	//=> [1 2 3 4 5 6 8]
	// fmt.Println(left)
	//=> 3
	//=> 1
	//=> 0
	//=> 1
	// fmt.Println(right)
	//=> 6
	//=> 2
	//=> 2
	//=> 1
	list[left], list[right] = list[right], list[left]
	// fmt.Println(list)
	//=> [3 1 2 4 8 6 5]
	//=> [1 2 3]
	//=> [5 6 8]
	//=> [6 8]
	//=> [1 2 3 4 5 6 8]

	quickSort(list[:left])
	quickSort(list[left+1:])

	return list
}
