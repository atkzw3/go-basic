package main

import "fmt"

func main() {
	var fl64 float64 = 2.5
	fmt.Println(fl64)

	// 暗黙的定義では、int型とは違く必ず float64 となる
	fl := 5.55
	fmt.Println(fl)
	fmt.Printf("%T\n", fl)

	// 型変換
	fl32 := float32(fl)
	fmt.Printf("%T\n", fl32)

	// 計算注意
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // +Inf でプラスの無限となる

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf でマイナスの無限

	nan := zero / zero
	fmt.Println(nan) // NaN
}
