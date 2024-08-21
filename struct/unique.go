package main

import "fmt"

// MyInt 型を作れる
type MyInt int

func main() {
	var i MyInt
	fmt.Println(i)
	fmt.Printf("%T\n", i) // main.MyInt

	//r := 100
	// 型が違うので計算することができない
	//fmt.Println(r * i)

	i.Greeting()
}

func (mi MyInt) Greeting() {
	fmt.Println("hello")
}
