package main

import "fmt"

//https://go.dev/ref/spec#Channel_types

// 複数のゴルーチン間でデータの受け渡しをするためのデータ構造
// データの送受信を行う

func main() {
	// 双方向
	var ch1 chan int
	fmt.Println(ch1) // <nil>

	// make関数を使うことで書き込み読み込みをすることができる
	ch1 = make(chan int)
	fmt.Println(ch1)

	// ==== 容量を設定することが可能 ====
	ch1a := make(chan int)
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch1a))

	ch1b := make(chan int, 5)
	fmt.Println(cap(ch1b)) //5

	// データをチャネルに送信
	ch1b <- 12345
	fmt.Println(len(ch1b)) // データは1つあることを確認

	ch1b <- 678
	fmt.Println(len(ch1b)) // 2つあること確認

	// 今回はキャパシティを5としているが、容量を超えた場合は、deadlockエラー出る

	// チャネルからデータを受信
	// 送信されたものから取得する特性
	i := <-ch1b
	fmt.Println(i)         // 12345
	fmt.Println(len(ch1b)) // 1
	i2 := <-ch1b
	fmt.Println(i2) // 678
	// 受信するとデータ消える
	fmt.Println(len(ch1b)) //0

	// === 明示的に受信や送信専用のチャネルを作成する場合もある ===
	// 受信
	var ch2 <-chan int
	fmt.Println(ch2) //<nil>
	// 送信
	var ch3 chan<- int
	fmt.Println(ch3) //<nil>
}
