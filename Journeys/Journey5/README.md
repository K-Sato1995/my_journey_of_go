# Maps
 A `map` is an unordered collection of key-value pairs, where each key is unique. `Maps` in Go are sometimes called `hashes` or `dictionaries` in other programming languages.

### Declaring a map
 You can define a `map` using the following syntax.

```go
① make(map[key_type]value_type)
② map[key_type]value_type { key1: value1, key2: value2, ......., keyX: valueX}
```

### ① make(map[key_type]value_type)
  You can initialize a map using the built-in `make()` function. The `make()` function returns a `map` of the given type, initialized and ready for use. `Keys` and corresponding `values`  can be added to a map like the code below.

```go
package main

import "fmt"

func main() {
  var map1 = make(map[int]string)

  fmt.Println(map1) //=> map[]

  //Add keys and values to 'map1'
  map1[1] = "Go"
  map1[2] = "Ruby"
  fmt.Println(map1) //=> map[1:Go 2:Ruby]
}
````

### ② map[key_type]value_type { key1: value1, key2: value2, ......., keyX: valueX}  
  By using the `map literal`, you can initialize a `map` with some initial data.

```go
package main

import "fmt"

func main() {
  var map2 = map[int]string { 1: "Go", 2:"Ruby"}

  fmt.Println(map2) //=> map[1:Go 2:Ruby]
}
```

### Nil maps
 If you declare a `map` with the `map literal` without initial data, it would generate a `nil-map`. Needless to say, a `nil-map` does not contain any data. Moreover, any attempt to add any data to a `nil-map` causes a runtime error.  

```go
package main

import "fmt"

func main() {
  var nil_map map[int]string

  if nil_map == nil {
    fmt.Println("nil") //=> nil
  }

  nil_map[1] = "GO" //=>  assignment to entry in nil map
}
```
### Modyfing a map
 You can add data to a `map` and modify the data of a `map` like the folloing code.

```go
package main

import "fmt"

func main(){
  var map_ex = make(map[int]string)


  //add keys and  values to a map
  map_ex[1] = "GO"
  map_ex[2] = "Ruby"

  fmt.Println(map_ex) //=> map[1:GO 2:Ruby]

  //modify the data of a map
  map_ex[1] = "Python"
  map_ex[2] = "Java"

  fmt.Println(map_ex) //=> map[1:Python 2:Java]
}
```

### Retrieving values from a map
 You can retrieve the value assigned to a key in a map using the syntax `map[key]`.

```go
package main

import "fmt"

func main(){
  var map_ex = make(map[int]string)

  map_ex[1] = "GO"
  map_ex[2] = "Ruby"

  fmt.Println(map_ex[1]) //=> GO
```

### Checking the existence of a key in a map
 When you retrieve the value assigned to a given key using the syntax `map[key]`, it returns an additional boolean value as well. It returns `true` if the key exists in the map and returns `false` if it does not exist in the map.

```go
package main

import "fmt"

func main(){
  var map_ex = make(map[int]string)

  map_ex[1] = "GO"
  map_ex[2] = "Ruby"

  //Trying to retrieve a key that exists in map_ex.
  lang, ok := map_ex[1]

 //Trying to retrieve a key that does not exist in map_ex.
  lang2, ok2 := map_ex[4]

  fmt.Println(lang, ok) //=> GO true
  fmt.Println(lang2, ok2) //=>  false
}
```

 If you just want to check for the existence of a key without retrieving the value associated with that key, then you can use an `_ (underscore)` in place of the first value.

```go
package main

import "fmt"

func main(){
  var map_ex = make(map[int]string)

  map_ex[1] = "GO"
  map_ex[2] = "Ruby"

  _, ok := map_ex[1]
  _, ok2 := map_ex[4]

  fmt.Println(ok) //=> true
  fmt.Println(ok2) //=> false
}
```

### Deleting a key from a map
You can delete a key from a map using the built-in `delete(map, key)` function. The `delete()` function does not do anything if the key does not exist in the map.

```go
package main

import "fmt"

func main(){
  var map_ex = make(map[int]string)

  map_ex[1] = "GO"
  map_ex[2] = "Ruby"

  delete(map_ex, 1)
  fmt.Println(map_ex) //=> map[2:Ruby]
}
```

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
