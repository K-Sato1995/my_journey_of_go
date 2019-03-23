package main

import "fmt"

func main() {
  intBasic()
  intValue()
  typeAssertions()
  typeSwitch(21) //=> Twice 21 is 42
  typeSwitch("Hello World") //=> "Hello World" is 11 bytes long
  typeSwitch(true)//=> I don't know about type bool!
  bob := Person{"Bob"}
  fmt.Println(bob.intro()) //=> Hello I'm Bob
}

func intBasic(){
  var intface interface {}
  //A variable of empty interface type can hold values of any type since every type implements at least zero methods.
  intface = 1
  fmt.Println(intface) //=> 1

  intface = "String"
  fmt.Println(intface) //=> String

  intface = []string{"Go", "Ruby", "JS"}
  fmt.Println(intface) //=> [Go Ruby JS]
}

func intValue(){
  var intface interface {}
  intface = 1
  fmt.Printf("%v %T\n", intface, intface) //=>1 int

  intface = "string"
  fmt.Printf("%v %T\n", intface, intface) //=> string string

  intface = []string{"Go", "Ruby", "JS"}
  fmt.Printf("%v %T\n", intface, intface) //=> [Go Ruby JS] []string
}


func typeAssertions(){
  var intface interface {} = "Hello World"

  t := intface.(string)
  fmt.Println(t) //=> Hello World

  t2, ok := intface.(string)
  fmt.Println(t2, ok) //=> Hello World true

  t3, ok := intface.(float64)
  fmt.Println(t3, ok) //=> 0 false
}


func typeSwitch(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

type People interface {
 intro()
}

type Person struct {
  name string
}

func (rarg Person) intro() string{
 return "Hello" + " " + "I'm" + " " + rarg.name
}
