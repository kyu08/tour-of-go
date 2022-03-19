# module
- 大文字スタートの名前は外部のパッケージから参照できる(exported name)
- 複数の返り値を返す関数の書き方はこう
  ```go
  // 定義側
  func swap(x,y string) (string, string) {
    return y, x
  }

  // 呼び出し側
func main () {
  a, b := swap("1", "2")

}
  ```

- 初期化しなくてもコンパイル通るんかーーーい
```go
package main

import (
	"fmt"
)

var c, python bool

func main() {
	var i int
	fmt.Println(i, c, python) // 0, false, false
}
```

- 変数宣言
  - varによる宣言
    - 
  - := による暗黙的な宣言
    - 関数の中では var の代わりに := を使って暗黙的な型宣言ができる
    - 関数の外ではキーワードで始まる宣言(func, var など)が必要で、 := による暗黙的な宣言は利用できない

- var でまとめて変数宣言してあると読みやすそう
```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

- Basic types
  - int は実行環境の OS や CPU の実装系に依存する
  - uint: 符号なし整数
  complex64(complex128): 複素数型

- 'a' が int32 ...?
  - go でのシングルクオートは `rune` という型として扱われる。何か文字を渡すと Unicode Code Pointを返す(文字のID)
```go
func main() {
	v := 'a'
	fmt.Printf("v is type of  %T\n", v) // v is type of int32
}
```

- const
  - 定数は `const` キーワードを使って宣言する
  - 文字, 文字列, boolean, 数値 のみが宣言できる
  - 定数の宣言に := を使うことはできない
  - 数値の定数は高精度な値
    - 型が定義されていない場合、その状況によって必要な型をとる
