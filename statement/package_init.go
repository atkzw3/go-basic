package main

import "fmt"

// https://go.dev/ref/spec#Package_initialization

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	// main関数が先ではなく、init()が先に読み込みされる
	// 複数設定可能 初期化したい時に使用する
	fmt.Println("main")
}
