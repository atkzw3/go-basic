package main

import "fmt"

// https://go.dev/ref/spec#Array_types
// 特徴： 要素数を後から変更できない
// スライス型であれば変更可能
func main() {
	var arr1 [3]int
	fmt.Println(arr1) // [0 0 0]
	// 型が要素数とtype合わせて型となる
	fmt.Printf("%T\n", arr1) // [3]int

	// 要素数3に対して、初期値は2つ。この場合 3つめは空文字が入る
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) // [A B ]

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("%T\n", arr3) // [3]int

	// 要素数の省略
	arr4 := [...]string{"A", "B"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4) // [2]string

	// 値の取り出し
	fmt.Println(arr4[0]) //A

	// 値の更新
	arr4[0] = "aaaaaaaa"
	fmt.Println(arr4) // [aaaaaaaa B]

	// 要素数をカウントする
	arrCount := len(arr3)
	fmt.Println(arrCount) //3
}
