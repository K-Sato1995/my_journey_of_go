# Overview
 In this post, I'll mainly explain what `pointers` and `structures` are.  `Pointers` and `Structures` can be very hard concepts to grasp for people who previously only learned script languages like `Ruby`, `Python` and so on.

# Pointers

### How to declare a pointer
 In brief, a `pointer` is a value which points to the memory address of another value. A `pointer` can be declared using the `*` followed by the type of the stored value in Go.      
The variable `pointer1` is a pointer to an `int` and `pointer2` is a pointer to a `string` in the following code.

```go
package main

import "fmt"

func main() {
  var pointer1 *int
  var pointer2 *string

  fmt.Println(pointer1, pointer2) //=> nil nil // A pointer's zero value is `nil`
}
```

### How to retrieve the memory address of a variable
 In Go, you can use the `&` operator to retrieve the memory address of a variable. In the following code, `&str` represents the memory address of the variable `str`.

```go
package main

import "fmt"

func main() {
  var str = "Go"
  fmt.Println(&str) //=> 0xc42000e1e0
}
```

### How to use pointers
 You can assign the memory address of a variable to a `pointer` just like the code below.

```go
package main

import "fmt"

func main() {
  var pointer1 *int
  var pointer2 *string

  var str = "Go"
  var num = 10

  pointer1 = &num
  pointer2 = &str

  fmt.Println(pointer1, pointer2) //=> 0xc420016078 0xc42000e1e0
}
```

 You can also declare a `pointer` more succinctly like the following code.

```go
package main

import "fmt"

func main() {
  var num2 = 20
  pointer3 := &num2

  fmt.Println(pointer3) //=> 0xc420016080
}
```

 `pointer1` and `pointer3` can only store the memory address of an `int` and `pointer2` can only store the memory address of a `string` respectively.

```go
package main

import "fmt"

func main() {
	var pointer1 *int
	var str = "Go"

	pointer1 = &str //=> cannot use &str (type *string) as type *int in assignment
}
```

### Accessing the value of a variable through its pointer
 The `*` operator denotes the pointer's underlying value. By using this feature, you can set, read and change the value of a variable through its pointer.

```go
package main

import "fmt"

func main() {
  var pointer1 *int
  var num = 10

  pointer1 = &num

  fmt.Println(*pointer1) //=> 10 //read var1 through the pointer.
  *pointer1 = 20 //change the value of var1 through the pointer.
  fmt.Println(num) //=> 20 //the new value of num
}
```

# Structs
 The concept of `structs` in Go is kind of similar to the concept of `classes` in object-oriented programming languages. A `struct` is a typed collection of fields. They’re useful for grouping data together to form records.

### How to declare a struct
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

### Different ways of Struct Instantiation
 There are multiple ways to instantiate a `struct`, I'll demonstrate the following 4 methods here.

####  The `var` keyword and Dot Notation
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

####  The `var` keyword and `:=` operator
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

#### Using the `new` keyword
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

### Structs and Pointers
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
# Methods
 As I mentioned earlier, Go does not have `classes`. However, you can define methods on struct types.    
 A `method` is a function with a special `receiver` argument.

### How to define a method
You can declare a method like the following code snippet.

```go
func(receiver receiver_type) name_of_function(arg) type_of_return_value {
   //content of the function
}
```

 In the example below, `intro` method has a `receiver` of type `person` named named `one_person`.

```go
package main

import "fmt"

type person struct {
    first_name string
    age int
}

func(one_person person) intro(arg string) string {
    return "Hello I'm" + " " + one_person.first_name + " " + arg     
}

func main() {
    mike := person{ "Mike", 20 }
    fmt.Println(mike.intro("!")) //=>Hello I'm Mike !
}
```

### Pointer receivers
 You can declare `methods` with pointer receivers. This means the `receiver_type` has the literal syntax `*Type` for some type `Type`.    
 There are several differences between `pointer receivers` and normal `value reveivers`.  One circumstance you should use `pointer receivers` over `value receivers` is when you want to change the state of the `receiver` in a method.    
 With `pointer receivers`, you can modify(read/write) the `receiver` in a method just like the code below.

```go
package main

import "fmt"

type Num struct {
    x, y int
}

func (p Num) stay() {
    p.x = 10
    p.y = 20
}

func (p *Num) modify() {
    p.x = 10
    p.y = 20
}

func main() {
    num_one := Num {0, 0}
    num_one.stay()

    fmt.Println(num_one) //=> {0 0}

    num_one.modify()

    fmt.Println(num_one) //=> {10 20}
}
```
 As you can see in the code above, The method `stay()` does not change the values of `Num` struct. On the other hand, The method `modify()` changes the values of `Num` struct.    
 If you want to know more about the differences of `pointer receivers` from `value receivers`, you can check [here](https://flaviocopes.com/golang-methods-receivers/)

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
