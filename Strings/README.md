# Introduction

I listed things that would come in handy when you are using string.

# (1) Multiline strings

Creating a multiline string in Go is actually incredibly easy. Simply use the backtick (\`) character when declaring or assigning your string value.

```go
str := `This is a
multiline
string.`
```

# (2) Efficient concatenation (Array => String)

You can concatenate an array of strings using `strings.Join`.

```go
package main

import (
  "fmt"
	"strings"
)

func main() {
  str := []string { "Go", "Ruby"}

  fmt.Println(strings.Join(str, "/")) //=> Go/Ruby
}
```

# (3) Convert ints (or any data type) into strings

You should look to use packages like strconv or functions like fmt.Sprintf. For example, here is an example using strconv.Itoa to convert an integer into a string.

```go
package main

import (
  "fmt"
	"strings"
  "strconv"
)

func main() {
  i := 123
  array := []string { "A", "B" }
  fmt.Println(strconv.Itoa(i)) //=> 123
  fmt.Println(returnString(array)) //=> The array becomes a string. [A B]
}


func returnString(arr []string) string{
  return fmt.Sprintf("The array becomes a string. %v", arr)
}
```

# (4) String into a slice

Use strings.Split on it.

```go
fmt.Printf("%#v\n", strings.Split("abc", "")) //=> []string{"t", "e", "s", "t"}
```

# (5) ToLower/ToUpper

```go
fmt.Println(strings.ToLower("STR")) //=> str
fmt.Println(strings.ToUpper("str")) ///=> STR
```

# References

[6 Tips for Using Strings in Go](https://www.calhoun.io/6-tips-for-using-strings-in-go/)
