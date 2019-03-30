package main

import "fmt"

func main() {
	// Declare a slice

	// (1) var name[]Type
	var slice1 []string
	fmt.Printf("%v\n", slice1) //=> []

	// (2) var name[]Type = []Type { ele1, ele2, elen.... }
	var slice2 []string = []string{"Go", "Ruby"}
	fmt.Printf("%v\n", slice2) //=> [Go Ruby]

	// (3) var name = array[start:end]
	// A slice is a segment of an array. That means we can create a slice from an array.
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice3 := arr[0:2] //=> [1,2]
	slice4 := arr[0:]  //=> [1 2 3 4 5 6]
	slice5 := arr[:5]  //=> [1 2 3 4 5]
	slice6 := arr[:]   //=> [1 2 3 4 5 6]
	fmt.Printf("%v\n", slice3)
	fmt.Printf("%v\n", slice4)
	fmt.Printf("%v\n", slice5)
	fmt.Printf("%v\n", slice6)

	// (4) name := make( []Type, len, cap)
	slice7 := make([]int, 2, 2)
	fmt.Printf("%v\n", slice7) //=> [0 0]

	// Modifying an array through a slice
	arr2 := [2]string{"Go", "Ruby"}
	slice8 := arr2[:]

	slice8[0] = "Javascript"
	fmt.Printf("%v\n", arr2) //=> [Javascript Ruby]

	// Length and Capacity of a Slice
	arr3 := [4]int{1, 2, 3, 4}
	slice9 := arr3[:1]
	fmt.Printf("The length is %v.The capacity is %v\n", len(slice9), cap(slice9))
	//=> The length is 1 .The capacity is 4

	// Appending to a slice
	slice10 := []int{1}
	slice10 = append(slice10, 2, 3)
	new_slice := append(slice10, 4, 5)
	fmt.Printf("%v\n", slice10)   //=> [1 2 3]
	fmt.Printf("%v\n", new_slice) //=> [1 2 3 4 5]

	// Nil slices
	var nil_slice []int

	if nil_slice == nil {
		fmt.Println("nil") //=> nil
	}
}
