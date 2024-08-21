package main

import "fmt"

// https://go.dev/ref/spec#Pointer_types

// ポインタとは、値型に分類されるデータ構造(基本型、参照型、構造体）のメモリ上のアドレスと型の情報のこと
// データ構造を間接的に参照、操作ができる
// 主に値型の操作に使われ、参照型は元からその機能を含んでいる為、基本的には不要
func main() {
	var n int = 100
	fmt.Println(n) //100
	// 変数の前にアンドをつける
	fmt.Println(&n) // 0xc000012078 (変わるので注意)

	Double(n)
	// 2倍にはならない
	// 値型は引数で渡したさいはコピーが行われるので更新されない
	fmt.Println(n) // 100

	// ポインタ型
	var p *int = &n
	fmt.Println(p) // 0xc000012078 (変わるので注意)

	// 値を取得する際は変数前に * をつける
	fmt.Println(*p) // 100

	*p = 300
	fmt.Println(*p) //300

	n = 200
	// nのアドレスを渡しているため更新される
	fmt.Println(*p) //200

	Double2(&n)
	fmt.Println(n) // 400

	// slice などの参照型は、引数として渡す場合はコピーではないので更新される
	var sl []int = []int{1, 2, 3}
	Double3(sl)
	fmt.Println(sl) //[2 4 5]

}

func Double(i int) int {
	i = i * 2
	return i
}

func Double2(i *int) int {
	*i = *i * 2
	return *i
}

func Double3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}
