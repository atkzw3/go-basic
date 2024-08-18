package main

import "fmt"

func main() {
	// https://go.dev/ref/spec#Break_statements
aiu:
	for {
		for {
			for {
				fmt.Println("START")
				break aiu
			}
			fmt.Println("END")
		}
		fmt.Println("終わりだよ")
	}
	fmt.Println("終わり〜")

hahahah:
	for i := 0; i < 5; i++ {
		fmt.Println("======== i開始 ========")
		fmt.Println("i = ", i)
		for j := 1; j < 5; j++ {
			fmt.Println("======== J 開始 ========")
			fmt.Println("j = ", j)

			// 4の時は処理を最初からしたい場合
			if j > 3 {
				continue hahahah
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("======== J 終了 ========")
		fmt.Println("処理しない")
	}
}
