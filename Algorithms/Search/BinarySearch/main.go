package main

import "fmt"

func main() {
	var sortedList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result, count := binarySearch(sortedList, 7)

	fmt.Printf("The index of the target is %v\n", result)
	fmt.Printf("It took %v times to search it.\n", count)
}

// This is referenced from the answer on stackOverFlow.
// https://stackoverflow.com/questions/43073681/golang-binary-search

func binarySearch(list []int, target int) (result int, count int) {
	var middleIndex = len(list) / 2
	fmt.Println(list)
	//=> [1 2 3 4 5 6 7 8 9 10]
	//=> [7 8 9 10]
	//=> [7 8]
	//=> [7]

	switch {
	case len(list) == 0:
		result = -1 // not found
	case list[middleIndex] > target:
		result, count = binarySearch(list[:middleIndex], target)
		// list[:middleIndex] means you are getting a list of the smaller values in the array
	case list[middleIndex] < target:
		result, count = binarySearch(list[middleIndex+1:], target)
		// list[middleIndex+1:] means you are getting a list of the larger values in the array
		result += middleIndex + 1
	default:
		result = middleIndex
	}
	count++
	return
}
