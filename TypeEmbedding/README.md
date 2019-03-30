# Type Embedding
Go does not provide the typical, type-driven notion of sub-classing, but it does have the ability to “borrow” pieces of an implementation by embedding types within a struct or interface.

```go
package main

import "fmt"

type Person struct {
   first_name string
}

func (a Person) name() string{ //Person typed method
    return a.first_name
}

type User struct {
     Person
}

func (a User) name() string { //User typed method
    return a.first_name
}

func main(){
  bob := Person{"Bob"}
  mike := User{}
  mike.first_name = "Mike"

  fmt.Println(bob.name()) //=> Bob
  fmt.Println(mike.name()) //=> Mike
}
```
