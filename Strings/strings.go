package main

import (
	"fmt"
	"strconv"
	"strings"
  "log"
  "regexp"
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

  // log
  log.Println("Log test")
  log.Printf("Log %v", array)


  // regexp
  const sample = `Test - test`
  var re = regexp.MustCompile(`\b|\W`)
  s := re.ReplaceAllString(sample, ``)
  fmt.Println(s) //=> Test test
}

func returnString(arr []string) string {
	return fmt.Sprintf("The array becomes a string. %v", arr)
}
