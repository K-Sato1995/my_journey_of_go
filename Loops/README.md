# For
 `For` is Go's only looping construct. Basic `for loop` consists of three parts that are listed below.

- `init statement`: It is executed before the first iteration. (It is often used to define a variable that can only be used in the scope of the `for` statement.)

- `condition`: It is evaluated before every iteration. (If it's true, the loop body executes. otherwise, the loop terminates.)

- `post statement`: It is executed at the end of every iteration.

 The basic syntax of Go's `for loop` looks like this.

```go
  for init statement; condition; post statement {
    //content of the loop
  }
```

The loop terminates the iteration once the condition evaluates to `false`.

```go
package main

import "fmt"

func main(){
  var num = 0

  for i := 0; i < 5; i++ {
      num += i
  }

  fmt.Println(num) //=> 10
}
```

`init` and `post` statements can be omitted. When those statements are omitted, Go's `for loop` behaves like the `while loop` of Javascript/Java/C.

```go
package main

import "fmt"

func main() {
	num := 1
	for ;num < 100; {
		num += num
	}
	fmt.Println(num) //=> 128
}
```

 If you omit the `condition`, it loops infinitely.

```go
package main

func main() {
	for {
	}
}
```
