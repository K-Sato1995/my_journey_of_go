package main

import "fmt"

func main() {
	m := map[int]string{1: "Go", 2: "Ruby", 3: "JavaScript"}
	m[4] = "TypeScript"

	fmt.Println(m) //=> map[1:Go 2:Ruby 3:JavaScript 4:TypeScript]

	m2 := map[int]map[int]string{1: {1: "Go"}, 2: {2: "Ruby"}}

	fmt.Println(m2) //=> map[1:map[1:Go] 2:map[2:Ruby]]
}
