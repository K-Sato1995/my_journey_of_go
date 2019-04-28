package main

import (
	"./first"
	"./second"
	"./zoo"
	"fmt"
)

func main() {
	// Printf
	var number = 10
	var str = "string"
	fmt.Printf("The number is %d. The str is %s\n", number, str)

	// Export, Import
	first.Exported() //=> This is Exported fron first.go
	second.Test()    //=> This is gonna be exported
	second.Bar()     //=> This is also exported
	fmt.Printf("%s\n", zoo.Lion()) //=> GAH
	fmt.Printf("%s\n", zoo.Dog()) //=> HAHA
}
