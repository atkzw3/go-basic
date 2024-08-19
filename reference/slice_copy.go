package main

import "fmt"

func main() {
	sl := []int{2, 3, 5, 7, 11, 13}
	// メモリが同じになる
	sl2 := sl
	// sl2だけではなく、slも更新されてしまう
	// 参照型のみなので注意
	// 参照型 => slice map interface chanel
	sl2[0] = 100
	fmt.Println(sl, sl2) // [100 3 5 7 11 13] [100 3 5 7 11 13]

	i := 1
	i2 := i
	i2 = 10000
	fmt.Println(i, i2) //1 10000

	sl3 := []int{2, 3, 5, 7, 11}
	sl4 := make([]int, 5, 10)
	n := copy(sl4, sl3)
	sl4[0] = 5000000
	sl4[1] = 1000000
	// nはコピーに成功した数
	fmt.Println(sl3, sl4, n) // [2 3 5 7 11] [5000000 1000000 5 7 11] 5
}
