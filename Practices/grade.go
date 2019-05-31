package main

import "fmt"

func main() {
	var scores = []int{50, 70, 80, 90}

	for _, v := range scores {
		fmt.Printf("Your grade is %v\n", checkScore(v))
	}

	/*
		Your grade is F
		Your grade is C
		Your grade is B
		Your grade is A
	*/
}

func checkScore(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else {
		return "F"
	}
}
