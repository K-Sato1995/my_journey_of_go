package main

import "fmt"

var COINS = []int{10, 50, 100, 500}
var CNUM = []int{0, 0, 0, 0}

func main() {
	money := 500
	price := 320

	change := (money - price)

	i := 2

	for 0 <= i {
		fmt.Println(COINS[i])
		CNUM[i] = change / COINS[i]
		change = change % COINS[i]
		i -= 1
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("You need %d %d\n", CNUM[i], COINS[i])
	}
}
