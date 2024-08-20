package main

import (
	"fmt"
	"time"
)

// https://go.dev/ref/spec#Close

func receiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			fmt.Println("break")
			break
		}
		fmt.Println("receive "+name, i)
	}
}

func main() {
	ch1 := make(chan int)
	close(ch1)

	// closeしたchannelに値を送るとruntime error が発生する
	// panic: send on closed channel
	//ch1 <- 1234

	// === 送信はできないが、closeしたchannelから受信は可能 ===

	ch2 := make(chan int, 2)
	ch2 <- 1234567
	close(ch2)

	//fmt.Println(<-ch2) // 1234567
	i, ok := <-ch2
	// データがない かつ closeでfalseとなる
	fmt.Println(i, ok) // 1234567 true

	ch3 := make(chan int)
	go receiver("A", ch3)
	go receiver("B", ch3)
	go receiver("C", ch3)

	i2 := 0
	for i2 < 100 {
		ch3 <- i2
		i2++
	}

	close(ch3)
	time.Sleep(3 + time.Second)
}
