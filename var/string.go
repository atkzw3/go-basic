package main

import "fmt"

func main() {
	var s string = "hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si) // string

	// 複数行 または " を入れたい場合ははバッククォーテーション
	fmt.Println(`a 
	b
	c """"`)

	fmt.Println(s[0]) // byte数
}
