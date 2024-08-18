package main

import "fmt"

// https://go.dev/ref/spec#For_statements
func main() {
	// 無限
	//for {
	//	fmt.Println("loop")
	//}

	// 5の時に処理終了
	i := 0
	for {
		i++
		if i == 5 {
			break
		}
		fmt.Println("loop")
	}

	// 10以下までループする
	point := 0
	for point < 10 {
		point++
		fmt.Println(point)
	}

	// 上記同様の書き方
	for i := 0; i < 10; i++ {

		// 3はskip
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// 配列
	fmt.Println("配列")
	array := [3]int{99, 100, 101}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// 範囲式方法

	// i はindex v はvalue
	for i, v := range array {
		fmt.Println(i, v)
	}

	// indexが不要な場合は _
	for _, value := range array {
		fmt.Println(value)
	}

	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println("map")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
