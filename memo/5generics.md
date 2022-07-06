## 型パラメータ
型パラメータは関数の引数の前に[] で囲んで記述する

```go
func Index[T comparable](s []T, x T) int
```

`comparable` は`==`と `!=` 演算子が使える便利な制約である

## Generic types 
構造体の定義にも型パラメータを用いることができる

```go
type List[T any] struct {
  next *List[T]
  val T
}
```

