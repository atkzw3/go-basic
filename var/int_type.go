package main

import "fmt"

// https://go.dev/ref/spec#Numerictypes
func main() {
	//int8        the set of all signed  8-bit integers (-128 to 127)
	//int16       the set of all signed 16-bit integers (-32768 to 32767)
	//int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	//int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// 数値によって型サイズは変更する

	// 型指定しない場合は、OS環境依存する
	var i int = 100
	fmt.Println(i)

	var i2 int8 = 127 // OK
	//var i2 int8 = 128 // NG
	fmt.Println(i2)

	// ========計算について========
	// 数値型でも同じ型ではないと計算できずエラーが出る
	var a int8 = 50
	var b int8 = 100
	var c int64 = 100
	fmt.Println(a + b)
	fmt.Println(c + 40)

	//fmt.Println(a + c) // NG

	// ========型変換========
	var i3 int64 = 300
	// 型表示
	fmt.Printf("%T\n", i3) // int64

	fmt.Printf("%T\n", int32(i2)) // int32

	i4 := int16(i2)
	fmt.Printf("%T\n", i4) //int16

}
