# Arrays
 In Go,  an `array` is a numbered sequence of elements of a single type with a fixed length(You can not resize `arrays` in Go). There are several ways to declare an `array` in Go. Here are some examples.

```go
① var name[num]Type
② var name[num]Type = [num]Type { ele1, ele2, elen.... }
③ name := [...] Type { ele1, ele2, elen... }

//name: The name of the array.
//num: The number of elements that array can contain.
//Type: The type of elements that the array contains.
```

 In the following code, I'll demonstrate how to declare an `arr` which is composed of 3 strings in the three ways that I presented above.

###  ① var name[num]Type

```go
package main

import "fmt"

func main(){
  var arr [3]string
  arr[0] = "Go"
  arr[1] = "Ruby"
  arr[2] = "Javascript"

  fmt.Println(arr) //=> [Go Ruby Javascript]
}
```

### ② var name[num]Type = [num]Type { ele1, ele2, elen.... }

```go
package main

import "fmt"

func main(){
  var arr [3]string = [3]string { "Go", "Ruby", "Javascript"}

  fmt.Println(arr) //=> [Go Ruby Javascript]
}
```

### ③ name := [...] Type { ele1, ele2, elen... }

```go
package main

import "fmt"

func main(){
  arr3 := [...]string { "Go", "Ruby", "Javascript" }

  fmt.Println(arr3) //=> [Go Ruby Javascript]
}
```
