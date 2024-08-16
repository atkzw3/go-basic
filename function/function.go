package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

// Divide 返り値複数ある場合は型に()をつける
func Divide(x, y int) (int, int) {
	return x / y, x % y
}

func Double(price int) (result int) {
	result = price * 2
	// 返り値の変数を指定することができる
	// 指定している場合はreturn は省略できる
	return
}

// NoReturn 返り値がない場合
func NoReturn() {
	fmt.Println("No return from function")
}

func main() {
	r := Plus(1, 1)
	fmt.Println(r)

	r2, r3 := Divide(2, 10)
	fmt.Println(r2, r3)

	r4 := Double(1000)
	fmt.Println(r4)

	NoReturn()
}
