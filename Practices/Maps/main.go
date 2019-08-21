package main

import "fmt"

func main() {
	m := map[int]string{1: "Go", 2: "Ruby", 3: "JavaScript"}
	m[4] = "TypeScript"
	v, ok := m[6]

	fmt.Println(m)  //=> map[1:Go 2:Ruby 3:JavaScript 4:TypeScript]
	fmt.Println(v)  //=> ""
	fmt.Println(ok) //=> false

	m2 := map[int]map[int]string{1: {1: "Go"}, 2: {2: "Ruby"}}

	fmt.Println(m2) //=> map[1:map[1:Go] 2:map[2:Ruby]]

	for k, v := range m {
		fmt.Println(k, v)
	}
	/*
		1 Go
		2 Ruby
		3 JavaScript
		4 TypeScript
	*/
}
