# Range
 The `range` keyword is used in `for loop` to iterate over items of an array, slice, channel or map. When you are ranging over an `array` or a `slice`, it returns the index and the element at that index. When you are ranging over a map, it returns the key and the key-value.

### Array

```go
package main

import "fmt"

var arry_ex  = [3]string { "Go", "Ruby", "Python" }

func main(){
  for index, value := range arry_ex {
    fmt.Println(index, value)
    //=> 0 Go
    //=> 1 Ruby
    //=> 2 Python
  }
}
```

### Slice

```go
package main

import "fmt"

var slice_ex = []string { "Go", "Ruby", "Python" }

  for index, ele := range slice_ex {
    fmt.Println(index, ele)
    //=> 0 Go
    //=> 1 Ruby
    //=> 2 Python
  }
}
```

### Map

```go
package main

import "fmt"

var map_ex = map[string]string{ "Name":"Sam", "Gender":"Male" }

  for key, value := range map_ex {
    fmt.Println(key, value)
    //=> Name Sam
    //=> Gender Male
  }
}
```
### Skipping the index, key or value
You can skip the `index` or `value` by assigning to `_`. If you only want the index, drop the `value` entirely.

### Index Only

```go
package main

import "fmt"

var arry_ex  = [3]string { "Go", "Ruby", "Python" }

func main(){
  for index, _ := range arry_ex {
    fmt.Println(index)
    //=> 0
    //=> 1
    //=> 2
  }
}
```


### Value Only

```go
package main

import "fmt"

var arry_ex  = [3]string { "Go", "Ruby", "Python" }

func main(){
  for _, value := range arry_ex {
    fmt.Println(value)
    //=> Go
    //=> Ruby
    //=> Python
  }
}
```
