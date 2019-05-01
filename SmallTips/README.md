# Go Small Tips

# 参照されないパッケージ

以下のように import する事で、コンパイルエラーを引き起こさないで参照できる。

```go
import _"fmt"
```

# まとめて変数や定数を定義

以下のようにまとめて変数や定数を定義する事ができる。

```go
var (
  n = 1
  s = "string"
)

const (
  t = true
  T = false
)
```

# 無名関数

```go
f := func(x, y int) int { return x + y }
f(3,5) //=> 8
```

# import で別名を使用する。

以下のようにする事で package 名を省略して書くことができる。

```go
import (
  f "fmt"
  . "test" // package名そのものを省略できる。
)

func main() {
  f.Println("hello")
  T // testから
}
```
