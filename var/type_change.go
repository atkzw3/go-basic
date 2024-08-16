package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Println(i2)
	fmt.Printf("i2 = %T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Println(i3) // 10 切り捨てで計算
	fmt.Printf("fl = %T\n", fl)

	// 文字列と数値変換

	var s string = "100"
	fmt.Println(s)
	fmt.Printf("s = %T\n", s)

	// 変換パッケージ 文字列から数値型へ変換 a to i
	// 2つ目のレスポンスは _ で対応。goでは定義した変数は必ず使用するルールがあるが、 _ にすることで使用しなくても良い
	i, _ = strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	// _ を使用しない場合は エラーハンドリングを追加する
	i, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	// === 数値 => 文字列 ===
	var i20 int = 20
	var s2 string = strconv.Itoa(i20)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	// === 文字列 => バイト配列 ===
	h := "hello world"
	b := []byte(h)
	fmt.Println(b) // [104 101 108 108 111 32 119 111 114 108 100]
	fmt.Printf("b = %T\n", b)

	h2 := string(b)
	fmt.Println(h2)
	fmt.Printf("h2 = %T\n", h2)
}
