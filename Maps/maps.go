package main

import "fmt"

func main() {
	// Define a map
	// (1) make(map[key_type]value_type)
	var map1 = make(map[int]string)
	fmt.Printf("%v\n", map1) //=> map[]

	// (2) map[key_type]value_type { key1: value1, key2: value2, ......., keyX: valueX}
	var map2 = map[string]string{"key1": "value", "key2": "value"}
	fmt.Printf("%v\n", map2) //=> map[key1:value key2:value]

	// Modyfing a map
	//Add keys and values to 'map1'
	map1[1] = "Go"
	map1[2] = "Ruby"
	map1[0] = "Ruby"

	fmt.Printf("%v\n", map1) //=> map[1:Go 2:Ruby 0:Ruby]

	// Change a value
	map1[1] = "Java"
	fmt.Printf("%v\n", map1) //=> map[0:Ruby 1:Java 2:Ruby]

	// Retrieve a value
	fmt.Printf("The value is %v.\n", map1[0]) //=> The value is Ruby.

	// Checking the existence of a key in a map
	// When you retrieve the value assigned to a given key using the syntax map[key], it returns an additional boolean value as well.
	map_ex := map[int]string{1: "Go", 2: "Ruby"}
	lang, ok := map_ex[1]
	lang2, ok2 := map_ex[3]

	fmt.Println(lang, ok)   //=> Go true
	fmt.Println(lang2, ok2) //=> false

	// just check the existance
	_, check := map_ex[1]
	_, check2 := map_ex[3]
	fmt.Println(check, check2) //=> true false

	// Deleting a key from a map
	map3 := map[string]string{"Key1": "value", "Key2": "Value"}
	delete(map3, "Key1")
	_, existance := map3["Key1"]
	fmt.Printf("%v\n", map3)      //=> map[Key2:Value]
	fmt.Printf("%v\n", existance) //=> false

	// Nil maps
	var nil_map map[string]string

	if nil_map == nil {
		fmt.Println("nil")
	}
	// nil_map[1] = "Go" //=>  assignment to entry in nil map
}
