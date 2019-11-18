package main

import "fmt"

// Defining a function
func greetings(name string) string {
	return "Hello" + " " + name
}

// Omission of the type of arguments
func multiple(num1, num2 int) int {
	return num1 * num2
}

// Multiple results
func reverse(str1, str2 string) (string, string) {
	return str2, str1
}

func main() {
	// Calling function1
	fmt.Println(greetings("John")) //=> Hello John

	// Calling function2
	fmt.Println(multiple(2, 3)) //=> 6

	// Calling function3
	a, b := reverse("Hello", "World")
	fmt.Println(a, b)                      //=> World Hello
	fmt.Println(reverse("Hello", "World")) //=> World Hello
}
