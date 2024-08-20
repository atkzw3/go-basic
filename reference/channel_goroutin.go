package main

import (
	"fmt"
	"time"
)

func receive(ch chan int, name string) {
	for {
		i := <-ch
		fmt.Println(i, name)
	}
}

// 複数のゴルーチンで1つ(複数)のchannelを共有することが多い
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go receive(ch1, "ch1")
	go receive(ch2, "ch2")

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i

		time.Sleep(1000 * time.Millisecond)
		i++
	}
	// 出力イメージ
	//0 ch1
	//0 ch2
	//1 ch2
	//1 ch1
	//2 ch2
	//2 ch1
	//3 ch2
	//3 ch1
	//4 ch2
	//4 ch1
	//5 ch2
	//5 ch1
	//6 ch2
	//6 ch1
	//7 ch2
	//7 ch1
	//8 ch2
	//8 ch1
	//9 ch2
	//9 ch1
	//10 ch2
	//10 ch1
}
