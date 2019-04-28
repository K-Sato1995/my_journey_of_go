# Effective Go

This is my summmary of [Effective Go](https://golang.org/doc/effective_go.html).

# Formatting

Use `go fmt` or `golint` to format your code.

# Commentary

Go provides block comments(`/* */`) and line commands(`//`).

Every package should have a package comment, a block comment preceding the package clause.

```go
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

If the package is simple, the package comment can be brief.

```go
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

# Names

Using `MixedCaps` is the convention in Go.

```go
// var example
var testVariable string

// function example
testFunction (){}
```

But if you want to export a function, you have to write it starting with a capital letter like the name below.

```go
TestFunction () {}
```

# Control structures

Since if and switch accept an initialization statement in Go, it's common to see one used to set up a local variable.

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

# Resources

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go package sources](https://golang.org/src/)
