## 型パラメータ
型パラメータは関数の引数の前に[] で囲んで記述する

```go
func Index[T comparable](s []T, x T) int
```

`comparable` は`==`と `!=` 演算子が使える便利な制約である

呼び出し側で型変数を定義する
が、推論可能な型な場合 language server が不要な旨を教えてくれるっぽい
```go
func main() {
	ints := map[string]int64{"john": 1, "taro": 2, "jiro": 3}
	fmt.Printf("sumInts: %d\n", sumIntsOrFloats[string, int64](ints))
}
```

## Generic types 
構造体の定義にも型パラメータを用いることができる

```go
type List[T any] struct {
  next *List[T]
  val T
}
```


