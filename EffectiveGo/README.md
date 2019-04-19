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

# Resources

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go package sources](https://golang.org/src/)
