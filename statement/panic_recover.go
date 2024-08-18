package main

import "fmt"

// https://go.dev/ref/spec#Handling_panics
// そもそもあまり使用されないので注意
func main() {

	// defer で遅延を行い、後の処理でpanic を発生させる
	// 停止せずにリカバリーできる

	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Recovered from panic:", x)
		}
	}()

	// panic は実行中のプログラムを停止する
	panic("panic recovered")
	fmt.Println("start")
}
