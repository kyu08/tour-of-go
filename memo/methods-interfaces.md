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


