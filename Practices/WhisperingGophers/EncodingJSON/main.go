package main

import (
	"encoding/json" // Package json implements encoding and decoding of JSON as defined in RFC 7159.
	"log"
	"os"
)

// Site has a title and an url
type Site struct {
	Title string
	URL   string
}

var sites = []Site{
	{"The Go programming language", "http://golang.org"},
	{"Google", "http://google.com"},
}

func main() {
	enc := json.NewEncoder(os.Stdout)
	for _, s := range sites {
		err := enc.Encode(s)
		if err != nil {
			log.Fatal(err)
		}
	}
}

/*
=>
{"Title":"The Go programming language","URL":"http://golang.org"}
{"Title":"Google","URL":"http://google.com"}
*/
