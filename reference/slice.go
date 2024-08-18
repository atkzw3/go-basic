package main

import "fmt"

// https://go.dev/ref/spec#Slice_types
func main() {
	// slice は可変長なので要素数は指定しなくてOK
	var slice []int
	fmt.Println(slice) // []

	var slice2 []int = []int{1, 2, 3}
	fmt.Println(slice2) // [1 2 3]

	slice3 := []string{"A", "B", "C"}
	fmt.Println(slice3) // [A B C]

	// make関数でslice作成
	// 2番目引数は要素数
	slice4 := make([]int, 3)
	fmt.Println(slice4)

	slice2[0] = 100
	fmt.Println(slice2)

	// 値の取り出し
	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice5)
	fmt.Println(slice5[0])                 //1
	fmt.Println(slice5[1:])                //[2 3 4 5]
	fmt.Println(slice5[:2])                //[1 2]
	fmt.Println(slice5[2:4])               //[3 4]
	fmt.Println(slice5[:])                 //[1 2 3 4 5]
	fmt.Println(slice5[len(slice5)-1])     // 5
	fmt.Println(slice5[1 : len(slice5)-1]) // [2 3 4]

	i := len(slice5) - 1
	fmt.Println(i) // 4
}
