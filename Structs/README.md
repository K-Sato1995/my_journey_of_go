# Structs

The concept of `structs` in Go is kind of similar to the concept of `classes` in object-oriented programming languages. A `struct` is a typed collection of fields. They’re useful for grouping data together to form records.

## How to declare a struct

You can declare a `struct` using the `type` and `struct` keywords.

```go
package main

import  "fmt"

type person struct {
  first_name string
  age int
}
```

In the example above, `person` struct contains `first_name` and `age` as its `fields`.

## Different ways of Struct Instantiation

There are multiple ways to instantiate a `struct`, I'll demonstrate the following 4 methods here.

### The `var` keyword and Dot Notation

A `struct` uses a `.` to access the values stored in fields.

```go
package main

import "fmt"

type person struct {
  first_name string
  age int
}

func main() {
  var mike person
  mike.first_name = "Mike"
  mike.age = 20

  fmt.Println(mike.first_name, mike.age) //=> Mike 20
}
```

### The `var` keyword and `:=` operator

The following 2 sets of code illustrates struct instantiation using `var` and `:=`.

```go
package main

import "fmt"

type person struct {
  first_name string
  age int
}

func main() {
  bob := person{ "Bob", 30 }

  fmt.Println(bob.first_name, bob.age) //=> Bob 30
}
```

You can also explicitly refer to a field and assign a value to it.

```go
package main

import "fmt"

type person struct {
  first_name string
  age int
}

func main() {
  sam := person{ age: 40, first_name: "Sam" }

  fmt.Println(sam.first_name, sam.age) //=> Sam 40
}
```

### Using the `new` keyword

The following code shows struct instantiation using the `new` keyword.

```go
package main

import "fmt"

type person struct {
  first_name string
  age int
}

func main() {
  jen := new(person)

  jen.first_name = "Jennifer"
  jen.age = 10

  fmt.Println(jen.first_name, jen.age) //=> Jennifer 10
}
```

## Structs and Pointers

`Struct fields` can be accessed through a struct pointer.

```go
package main

import "fmt"

type person struct {
  first_name string
  age int
}

func main() {
  var mike person

  s_pointer := &mike

  s_pointer.first_name = "Mike"
  s_pointer.age = 20
  fmt.Println(mike) //=> {Mike 20}
}
```

## Structure in another Structure

You can use a structure as a field of another structure.

```go
// Structure in another Structure
type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed Feed // Using Feed structure as a field
}
```
