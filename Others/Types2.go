package main

import "fmt"

type (
	intPair [2]int
	areaMap map[string][2]float64
)

func main() {
	pair := intPair{1, 2}
	area := areaMap{"Boston": {45.3, 67.5}}

	fmt.Println(pair, area)
}
