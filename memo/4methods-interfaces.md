# method
- 以下のようにすることで構造体・型にメソッドを追加することができる
- ※パッケージ内で定義している構造体・型に限る

```go
type User struct {
  age int
}

func (u User) add() User {
  u.age++
  return u
}

func main () {
  u := User{1}
  fmt.Printf("%#v", u.add) // main.User{age:2}
```

# pointer receiver
- レシーバそのものの値を変更したい場合は pointer receiver の形でメソッドを定義する

```go
type User struct {
  age int
}

func (*u User) add() User {
  u.age++
}

func main () {
  u := User{1}
  fmt.Printf("%#v", u) // main.User{age:1}
  u.add()
  // メソッドをポインタレシーバで定義したため u が変更されている
  fmt.Printf("%#v", u) // main.User{age:2}
```

# interface
インターフェースを実装することを明示的に宣言する必要はない
-> ある型がどんな interface なのか知りたくなった時にパッと調べることはできない？

# error
`error`型は `fmt.Stringer`に似た組み込みのインターフェースである
```go
type error interface {
  Error() string
}
```

## 疑問
暗黙的にインターフェイスを満たせるみたいだけどメリットあるのかな
満たしているインターフェイスを明示的に記述した方が可読性が高くて良さそうだと思った
https://go-tour-jp.appspot.com/methods/10



