package main

import "fmt"

func main() {
	sl := []int{1, 2, 3}
	fmt.Println(sl) // [1 2 3]

	// append
	sl = append(sl, 4, 5, 6)
	fmt.Println(sl) // [1 2 3 4 5 6]

	// make
	sl2 := make([]int, 3)
	fmt.Println(sl2) // [0 0 0]

	// len・cap
	// len => 要素数
	// cap => 容量
	fmt.Println(len(sl2), cap(sl2)) // 3 3

	sl3 := make([]int, 3, 5)
	fmt.Println(sl3)                // [0 0 0]
	fmt.Println(len(sl3), cap(sl3)) // 3 5
	// capはメモリなど気にする時に考慮する

	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	// 容量が 5から10の2倍になる
	// ======= 不用意に追加などしてメモリを多く使用しないこと =========
	fmt.Println(len(sl3), cap(sl3)) // 6 10
	// 過剰にメモリを確保してしまうと実行速度が落ちたりする。
	// 良質なパフォーマンスを実現するには、容量の管理も気にすること。
}
