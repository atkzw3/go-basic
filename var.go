package main

import "fmt"

var i5 int = 1000

func main() {
	// 明示的定義 (変数名 型 = 値)
	var i int = 100
	fmt.Println(i)

	var s string = "hello"
	fmt.Println(s)

	// 同型
	var t, f bool = true, false
	fmt.Println(t, f)

	// 別型
	var (
		i2 int    = 500
		s2 string = "hi!"
	)
	fmt.Println(i2, s2)

	// 値なしで変数に型のみ定義は初期値が自動で入る
	var i3 int    // 初期値 0
	var s3 string // 初期値 空文字
	fmt.Println(i3, s3)

	i3 = 1000
	fmt.Println(i3)
	s3 = "hey"
	fmt.Println(s3)

	// 値の更新
	fmt.Println(i)
	i = 150
	fmt.Println(i)

	// 暗黙的な定義(変数名 :=値)
	// 型推論 自動で値から型を定義してくれる
	// 関数外では使用できないので、使い分けはそこで判断
	// 基本は明示的で進める

	i4 := 400
	// 再定義はできない i4 :=900
	fmt.Println(i4)

	// 値の更新
	i4 = 450
	fmt.Println(i4)

	var d string = outer()
	fmt.Println(d)

	d2 := outer()
	fmt.Println(d2)
}

func outer() string {
	var ss string = "hello world"
	fmt.Println(ss)

	sss := "hello world!!!!"
	fmt.Println(sss)

	return sss
}
