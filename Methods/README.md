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
