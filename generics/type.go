package main

import "fmt"

// Number Numberには指定された型のみ
type Number interface {
	~int | ~int32 | ~int64
}

func Max[T Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

type MyInt2 int

func main() {
	fmt.Println(Max(100, 1))
	fmt.Println(Max(1, 2))

	var x, y MyInt2 = 1, 2
	// MyInt2 はint型でNumberを満たすが、エラーを発生させてしまう (仕様)
	// その場合はNumberの型の前に ~ をつける
	fmt.Println(Max[MyInt2](x, y))
}
