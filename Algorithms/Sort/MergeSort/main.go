package main

import "fmt"

func main() {
	var list = []int{3, 5, 1, 2, 8, 6, 4}
	fmt.Println(mergeSort(list)) //=> [1 2 3 4 5 6 8]
}

func mergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	mid := len(list) / 2 // Middle Index of the whole array.

	//list[:mid] means the left half of the array and list[mid:] means the right part of the array.
	return merge(mergeSort(list[:mid]), mergeSort(list[mid:]))
}

func merge(left, right []int) []int {
	// size = the size of the whole array.
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	// Make a new slice (also avoiding destroying the original, unsorted slice).
	// We then keep track of the left and right slices to be merged and just pick the next smallest value from the left or right slice to add to our new slice.
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
