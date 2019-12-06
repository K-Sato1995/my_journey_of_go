package main

import (
	"compress/gzip"   // Package gzip implements reading and writing of gzip format compressed files, as specified in RFC 1952.
	"encoding/base64" // Package base64 implements base64 encoding as specified by RFC 4648.
	"io"              // Package io provides basic interfaces to I/O primitives.
	"os"              // Package os provides a platform-independent interface to operating system functionality.
	"strings"         // Package strings implements simple functions to manipulate UTF-8 encoded strings.
)

// package io
// The io package provides fundamental I/O interfaces that are used throughout most Go code.
// type Writer interface {
//     Write(p []byte) (n int, err error)
// }
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }

func main() {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	r, _ = gzip.NewReader(r)
	io.Copy(os.Stdout, r)
}

const data = `
H4sIAAAJbogA/1SOO5KDQAxE8zlFZ5tQXGCjjfYIjoURoPKgcY0E57f4VZlQXf2e+r8yOYbMZJhoZWRxz3wkCVjeReETS0VHz5fBCzpxxg/PbfrT/gacCjbjeiRNOChaVkA9RAdR8eVEw4vxa0Dcs3Fe2ZqowpeqG79L995l3VaMBUV/02OS+B6kMWikwG51c8n5GnEPr11F2/QJAAD//z9IppsHAQAA
`
