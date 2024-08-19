package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	// sliceで渡すことができる
	fmt.Println(Sum(1, 2, 3, 4, 5))

	// sliceで渡すこともできる
	sl := []int{1, 2, 3, 4, 5}
	Sum(sl...)
}
