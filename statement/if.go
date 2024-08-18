package main

import "fmt"

// https://go.dev/ref/spec#If_statements
func main() {
	a := 3
	if a == 2 {
		fmt.Println(true)
	} else if a == 3 {
		fmt.Println("3だ")
	} else {
		fmt.Println(false)
	}

	// 前：簡易文 + 後 条件分岐
	if b := 100; b == 100 {
		fmt.Println("100だよ")
	}

	x := 0
	// 外側のx = 0 変数ではなく、x =2が優先される
	if x := 2; true {
		fmt.Println(x) //2
	}
	fmt.Println(x) //0
}
