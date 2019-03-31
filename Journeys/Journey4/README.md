# Introduction
 The concepts and use of `Arrays` and `Slices` have been quite confusing to me since I started learning Go. This is finally time for me to really understand how they work in Go and the differences between these two.

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

# Slices
 In Go, A `slice` is a segment of an array. `Slices` build on `arrays` and provide more power, flexibility, and convenience compared to `arrays`. Just like `arrays`, `Slices` are indexable and have a length. But unlike `arrays`, they can be resized.

### Declaring a Slice
 A `Slice` can be declared in the following ways. Unlike declaring an `array`, you don't have to specify the number of elements the `slice` can contain when you define it.

```go
① var name[]Type
② var name[]Type = []Type { ele1, ele2, elen.... }
③ name :=  Array[start:end]
④ name := make( []Type, len, cap)

//name: The name of the array.
//Type: The type of elements that the array contains.
//Array: An array.
//make: The built-in make function.
```
 I'll demonstrate how to declare a `slice`  using four ways of defining a `slice` I presented above in the following code.

### ① var name[]Type

```go
package main

import "fmt"

func main() {
  var slice1 []string
  var slice2 = []int { 1, 2, 3 }

  fmt.Println(slice1) //=> []
  fmt.Println(slice2) //=> [1 2 3]
}
```

### ② var name[]Type = []Type { ele1, ele2, elen.... }

```go
package main

import "fmt"

func main() {
  var slice2 []string = []string { "Go", "Ruby" }

  fmt.Println(slice2) //=> [Go Ruby]
}
```

### ③ name :=  Array[start:end]
 As I stated above, A `slice` is a segment of an `array`. That means we can create a slice from an array.  The table below explains how to control the elements a `slice` which is created from `Array` contains.

| 操作 | 意味 |
|:--|:--|
| Array[low:high] | From `low` to `high - 1`. |
| Array[low:] | From `low` to the last element. |
| Array[:high] | From the first elemet to `high - 1`. |
| Array[:] | From the first element to the last element. |

```go
package main

import "fmt"

func main() {
  arry := [6] int { 1, 2, 3, 4, 5, 6 }

  slice3 := arry[0:2]
  slice4 := arry[0:]
  slice5 := arry[:4]
  slice6 := arry[:]

  fmt.Println(slice3) //=> [1 2]
  fmt.Println(slice4) //=> [1 2 3 4 5 6]
  fmt.Println(slice5) //=> [1 2 3 4]
  fmt.Println(slice6) //=> [1 2 3 4 5 6]
}
```

### ④ name := make( []Type, len, cap)

```go
package main

import "fmt"

func main() {
  slice7 := make([]string, 2, 2)

  fmt.Println(slice7) //=> [ ]
}
```

### Modifying a slice
 As I mentioned several times, a `slice` is a segment of an `array` and it refers to an underlying array. Therefore, modifying the elements of a `slice` will also modify the corresponding elements in the referenced array.

```go
package main

import "fmt"

func main(){
  arry := [...]string { "Go", "Ruby", "Javascript" }

  slice := arry[:]

  fmt.Println(slice) //=> [Go Ruby Javacript]

  slice[0] = "Python"

  fmt.Println(slice) //=> [Python Ruby Javascript]
  fmt.Println(arry) //=> [Python Ruby Javascript]
}
```

### Length and Capacity of a Slice
 A `slice` has both a `length` and a `capacity`.

- The length of a slice is the number of elements it contains.
- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

You can find the length and capacity of a slice using the built-in functions `len()` and `cap() `.

```go
package main

import "fmt"

func main(){
  arry := [...]string { "Go", "Ruby", "Javascript" }

  slice := arry[:1]

  fmt.Println(len(slice), cap(slice)) //=> 1 3
}
```

### Appending to a slice
 It is common to append new elements to a slice, and so Go provides a built-in `append` function. It takes a slice and new elements that you want to append to the `slice`. It then returns a new slice containing all the elements from the given slice as well as the new elements.  
The following code shows the structure of `append` function.

```go
func append(slice []Type, new_elements) []Type

//Type: The type of elements that the array contains.
```

```go
package main

import "fmt"

func main(){
  arry := [...]string { "Go", "Ruby", "Javascript" }

  slice := arry[:]
  var new_slice = append(slice, "Java", "Swift", "C")

  fmt.Println(new_slice) //=> [[Go Ruby Javascript Java Swift C]
}
```

### Nil slices
The zero value of a slice is `nil`. A `nil slice` has a `length` and `capacity` of 0 and has no underlying array.

```go
package main

import "fmt"

func main(){
  var nil_slice []int

  if nil_slice == nil {
    fmt.Println("nil") //=> nil
  }
```
