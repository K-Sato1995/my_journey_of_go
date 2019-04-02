package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// (2) Array => String
	str := []string{"Go", "Ruby"}
	fmt.Println(strings.Join(str, "/")) //=> Go/Ruby

	// (3) int => string
	i := 123
	array := []string{"A", "B"}
	fmt.Println(strconv.Itoa(i))     //=> 123
	fmt.Println(returnString(array)) //=> The array becomes a string. [A B]
}

func returnString(arr []string) string {
	return fmt.Sprintf("The array becomes a string. %v", arr)
}
