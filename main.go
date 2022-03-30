package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	// [疑問1]なぜこれで無限ループにならずに "Hello, Reader!" をいい感じに読んでいけるのか
	// Read の実装を見ると `	r.i += int64(n)` とかしている。r.i の i は以下のように定義されており、 current reading index のことらしい。
	// type Reader struct {
	// s        string
	// i        int64 // current reading index
	// prevRune int   // index of previous rune; or < 0
	// }

	// [疑問2]なぜ 0 以外の n を Read から受け取れているのか
	// Read の実装を見ると正常系では何も値を返していないように見える
	// -> https://go-tour-jp.appspot.com/basics/7 これやで
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
