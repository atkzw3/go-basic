package main

import "fmt"

// CallBack 渡した関数の実行をしてくれる
func CallBack(f func()) {
	f()
}

func main() {
	// 無名関数を渡す
	CallBack(func() {
		fmt.Println("Hello World")
	})
}
