package main

import "fmt"

// https://go.dev/ref/spec#Switch_statements
func main() {
	// caseとのデータ型が一致しないとエラー出る
	n := 00
	switch n {
	case 1, 2:
		fmt.Println(n)
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("non")
	}

	// 上記と同じ 変数をswitchの後に定義している
	switch n := 2; n {
	case 1, 2:
		fmt.Println(n)
	// 列挙で1,2のように出している中で、条件式を書くことはできない
	//case n > 5:
	//	fmt.Println(n)
	case 3, 4:
		fmt.Println("3 or 4")
	}

	n1 := 6
	switch {
	case n1 < 0 && n < 4:
		fmt.Println("aaa")
	case n1 > 5:
		fmt.Println("eee")
	default:
		fmt.Println("rrr")
	}
}
