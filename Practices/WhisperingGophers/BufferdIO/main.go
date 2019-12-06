package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

// The bufio package implements buffered I/O. Its bufio.Scanner type wraps an io.Reader and provides a means to consume it by line (or using a specified "split function").
const input = `Input String`

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
