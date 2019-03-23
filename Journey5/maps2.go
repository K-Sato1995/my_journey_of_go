package main

import "fmt"

func main(){
  var map_ex = make(map[int]string)


  //add keys and  values to a map
  map_ex[1] = "GO"
  map_ex[2] = "Ruby"

  lang, ok := map_ex[1]
  lang2, ok2 := map_ex[4]

  fmt.Println(lang, ok) //=> GO true
  fmt.Println(lang2, ok2) //=>  false

  //Only checking the existance of a key.
  _, ok3 := map_ex[1]
  _, ok4 := map_ex[4]
  fmt.Println(ok3) //=> true
  fmt.Println(ok4) //=> false

  fmt.Println(map_ex) //=> map[1:GO 2:Ruby]
  fmt.Println(map_ex[1]) //=> GO

  //Deleting a key from a map
  delete(map_ex, 1)
  fmt.Println(map_ex) //=> map[2:Ruby]

  //modify the data of a map
  map_ex[1] = "Python"
  map_ex[2] = "Java"

  fmt.Println(map_ex) //=> map[1:Python 2:Java]
}
