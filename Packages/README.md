# Packages
 Every Go program consists of `packages`. Each package needs to be imported first to use its exported identifiers.  
 You can import `packages` like the code below.

 ```go
 package main

 import (
   "fmt" // Imported Package
   "math/rand" // Imported Package
 )

 func main(){
   fmt.Println("My journey of Go", rand.Intn(10)) //=>My journey of Go 1
 }

 ```

You can also import each package individually.

 ```go
 import "fmt"
 import "math/rand"
 ```

# Exported names

In Go, a name is exported if it begins with a capital letter. For instance,  `Pi` is a name that is exported from `'math'` package.

 ```go
 package main

 import "math"

 func main(){
   fmt.Println(math.Pi) //=> 3.141592653589793
 }
 ```
