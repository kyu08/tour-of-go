# switch 文 は条件を指定しなかった場合、case の評価結果が true であれば処理が実行される

```go
// 条件を指定するパターン
func main() {
  today := time.Now().Weekday()
     switch time.Saturday {
       case today + 0:
        println("today")
      case today + 1:
        println("tomorrow")
      default:
        println("too far away")
     }
}

// 条件を指定しないパターン
func main() {
  age := 20
     switch {
       case age < 20:
        println("miseinen")
      case age == 20:
        println("hatachi")
      default:
        println("21 ijou")
     }
}
```

# defer の実行順
defer は LIFO の順番で実行される
