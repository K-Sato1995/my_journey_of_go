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
	size, leftIndex, rightIndex := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	// Make a new slice (also avoiding destroying the original, unsorted slice).
	// We then keep track of the left and right slices to be merged and just pick the next smallest value from the left or right slice to add to our new slice.
	for k := 0; k < size; k++ {
		if leftIndex > len(left)-1 && rightIndex <= len(right)-1 {
			// leftIndex > len(left)-1 means left array is empty.
			// So put the smallest value from the right array and add 1 to the rightIndex.
			slice[k] = right[rightIndex]
			rightIndex++
		} else if rightIndex > len(right)-1 && leftIndex <= len(left)-1 {
			slice[k] = left[leftIndex]
			leftIndex++
		} else if left[leftIndex] < right[rightIndex] {
			slice[k] = left[leftIndex]
			leftIndex++ // To use the same value to compare from the right array.
		} else {
			slice[k] = right[rightIndex]
			rightIndex++ // To use the same value to compare from the left array.
		}
	}
	return slice
}
