package main

import (
	"compress/gzip"
	"log"
	"strings"
)

func main() {
	log.Println("Opeing gzip stream........")
	_, err := gzip.NewReader(strings.NewReader("not a gzp stream!"))

	// log.Fatal prints a message and exit the program printing a stack trace.
	if err != nil {
		log.Fatal(err)
	}

	log.Println("No error was detected")
}
