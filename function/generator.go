package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	// 値の保持は別となる
	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
}
