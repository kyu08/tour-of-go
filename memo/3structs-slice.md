# Pointers
- ポイントは値のメモリアドレスを指す
- 


[go print系関数の違い](https://qiita.com/taji-taji/items/77845ef744da7c88a6fe)
[go の名前空間について](https://ifritjp.github.io/documents/go/package/)


# slice
- make で slice を作成することができる
```go
s := make([]int, 7)
```

# map
- map も make で作成することできる
```go
type User struct {
  age int
}

m := make(map[string]User)
```
## 疑問
- 以下のように存在しないkey の value は int のゼロ値である0が返るけど、value が存在しないのか、 0 という値を持っているのかはどう判別するのがいいんだろう

```go
type User struct {
	age int
}

func main() {
	m := make(map[string]User)
	m["john"] = User{5}
	fmt.Printf("%+v\n", m["a"])
}

// 以下のようにするらしい
elem, ok := m[key] // ok: bool

```
