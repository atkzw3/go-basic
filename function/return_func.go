package main

import "fmt"

func Return() func() {
	return func() {
		fmt.Println("Return function")
	}
}

func main() {
	// 関数取得
	f := Return()
	// 関数実行
	f()
}
