package main

import "fmt"

func main() {
	var scoreList = [][]int{
		{34, 56, 78},
		{34, 67, 21},
		{90, 87, 79},

		// The Average score of class 0 is 56
		// The Average score of class 1 is 40
		// The Average score of class 2 is 85
	}

	for i, v := range scoreList {
		fmt.Printf("The Average score of class %d is %d\n", i, calAve(v))
	}
}

func calAve(scores []int) int {
	var total int

	for _, v := range scores {
		total += v
	}

	return total / len(scores)
}
