package main

import "fmt"

// Hello says Hello
func Hello(str string) string {
	if str == "" {
		return "Hello, World"
	}

	return fmt.Sprintf("Hello, %s", str)
}
