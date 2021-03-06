package main

import (
	"fmt"
	"log"
	"regexp"
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

	// (4) string => slice
	fmt.Printf("%#v\n", strings.Split("abc", "")) //=> []string{"t", "e", "s", "t"}

	// log
	log.Println("Log test")
	log.Printf("Log %v", array)

	// regexp
	const sample = `Test - test`
	var re = regexp.MustCompile(`\b|\W`)
	s := re.ReplaceAllString(sample, ``)
	fmt.Println(s) //=> Test test

	// ToLower/ToUpper
	fmt.Println(strings.ToLower("STR")) //=> str
	fmt.Println(strings.ToUpper("str")) ///=> STR
}

func returnString(arr []string) string {
	return fmt.Sprintf("The array becomes a string. %v", arr)
}
