package main

import "fmt"

// https://go.dev/ref/spec#Constants

// Pi 頭文字を大文字にすると他のパッケージでも呼び出せるようになる
const Pi = 3.14

// 複数は ()
const (
	URL  = "https://go.dev/ref/spec#Constants"
	Name = "golang docs"
)

const (
	A = 1
	B
	C
	D = "dddd"
	E
	F
)

var Big int = 123 + 1

const Big2 = 123 + 1

const (
	c0 = iota // 連続する整数の連番を返す
	c1
	c2
)

func main() {
	// 定数は関数ないでもOKだが、関数外で定義することが多い

	fmt.Println(Pi)
	fmt.Println(URL, Name)
	// 値を省略して入れることができる
	fmt.Println(A, B, C, D, E, F) // 1 1 1 dddd dddd dddd
	fmt.Println(Big)
	fmt.Println(Big2 + 1)
	fmt.Println(c0, c1, c2) // 0 1 2
}
