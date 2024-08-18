package main

import (
	"fmt"
	"os"
)

// https://go.dev/ref/spec#Defer_statements
// deferの使用例として、リソースの開放処理がある

func TestDefer() {
	defer fmt.Println("あいうえお") //表示順 2
	defer fmt.Println("かきくけこ") //表示順 1

	fmt.Println("最初に表示されるよ")
}

func RunDefer() {
	fmt.Println("run defer")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	// 無名関数の使い方
	// この場合は1・2・3の順番
	// 下のRunDefer関数より後に発火
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	// 表示順は3,2,1
	RunDefer()

	// osパッケージでtestファイル作成する
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("error = ", err)
	}
	// リソース開放処理を行う
	defer file.Close()

	file.Write([]byte("hello world"))
}
