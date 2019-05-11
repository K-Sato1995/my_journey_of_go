package main

import "fmt"

func main() {
	var list = []int{3, 5, 1, 2, 8, 6, 4}

	fmt.Println(insertionSort(list)) //=> [1 2 3 4 5 6 8]
}

func insertionSort(list []int) []int {
	var listLength = len(list)

	// assume list[0] is already sorted
	for i := 1; i < listLength; i++ {
		j := i
		for j > 0 {
			// list[j] is the current value
			// list[j-1] is the index of the sorted value that is one value before the current one.
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1] // swap the values if unsorted value is smaller that the sorted value.
			}
			j-- // Repeat it until the last sorted value.
		}
	}

	return list
}
