package main

import "fmt"

func main() {
	var list = []int{3, 5, 1, 2, 8, 6, 4}
	selectionSort(list)
	fmt.Println(list)
}

func selectionSort(list []int) []int {
	var listLen = len(list)

	// Loop1
	for i := 0; i < listLen; i++ {
		// fmt.Println(i)   // 0 1 2 3 4 5 6
		var minIndex = i // The index of the minimum value

		// Loop2
		for j := i; j < listLen; j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		// Swap the value on the far left with the minimum value.
		list[i], list[minIndex] = list[minIndex], list[i]
	}
	return list
}
