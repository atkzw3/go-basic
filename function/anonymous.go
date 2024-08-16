package main

import "fmt"

func main() {
	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i := f(3, 4)
	fmt.Println(i)

	// 引数を直接渡す場合は()を後ろにつけて渡す
	i2 := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(i2)

}
