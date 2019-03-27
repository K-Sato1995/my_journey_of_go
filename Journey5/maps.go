package main

import "fmt"

func main() {
	var map1 = make(map[int]string)
	var map2 = map[int]string{1: "Go", 2: "Ruby"}
	var nil_map map[int]string

	if nil_map == nil {
		fmt.Println("nil")
	}
	// nil_map[1] = "GO" //=>  assignment to entry in nil map

	fmt.Println(map1) //=> map[]

	//Add keys and values to 'map1'
	map1[1] = "Go"
	map1[2] = "Ruby"
	fmt.Println(map1) //=> map[1:Go 2:Ruby]
	fmt.Println(map2) //=> map[1:Go 2:Ruby]
}
