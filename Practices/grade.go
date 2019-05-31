package main

import "fmt"

func main() {
	var scores = []int{50, 70, 80, 90}

	for _, v := range scores {
		fmt.Println(checkScore(v))
	}
	/*
		F
		C
		B
		A
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
