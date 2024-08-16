package main

import "fmt"

// https://go.dev/ref/spec#Interface_types
func main() {
	// interfaceは型が変更できる
	// 計算・演算ロジックはできないので注意

	var x interface{}
	fmt.Println(x)        // 初期値は<nil> 値を持っていない
	fmt.Printf("%T\n", x) // <nil>

	// ==== interfaceは型が変更できる ====
	x = 1
	fmt.Println(x)
	fmt.Printf("%T\n", x) // int

	x = "aaa"
	fmt.Println(x)
	fmt.Printf("%T\n", x) // string
}
