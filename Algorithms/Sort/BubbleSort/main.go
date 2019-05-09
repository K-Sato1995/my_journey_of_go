package main

import "fmt"

func main() {
	var numbers []int = []int{5, 3, 4, 2, 1}
	fmt.Println("The original nummbers:", numbers)
	sweep(numbers)
	fmt.Println("After one sweep:", numbers)
	bubbleSort(numbers)
	fmt.Println("After bubbleSort:", numbers)
}

// The “loop” that ran the sweep N times, where N is the size of the list we are sorting.
func bubbleSort(numbers []int) {
	var N int = len(numbers)

	for i := 0; i < N; i++ {
		sweep(numbers)
	}
}

// Compare pairs of numbers in the array and swapped them if the first item in the pair was greater than the second.
func sweep(numbers []int) {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	// Looping unless secondIndex is smaller than the length of the slice
	for secondIndex < N {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			// Swap the numbers!
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}
}
