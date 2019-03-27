package main

import (
	"fmt"       // Imported Package
	"math"      // Imported Package
	"math/rand" // Imported Package
)

func main() {
	fmt.Printf("The value is %v\n", rand.Intn(10)) //=> The value is 1
	fmt.Printf("%v\n", math.Pi)                    //=> 3.141592653589793
}
