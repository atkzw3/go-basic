package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "ABC"

	// 複数のchannelを使用している際に、ch2は問題ないが、ch1に値がなくてdeadlockしてしまう場合がある
	// この問題を回避するためにはselectする
	//v1 := <-ch1  ここでdeadlock
	//v2 := <-ch2
	//fmt.Println("v1:", v1)
	//fmt.Println("v2:", v2)

	select {
	case v1 := <-ch1:
		fmt.Println("v1", v1)
	case v2 := <-ch2:
		fmt.Println("v2", v2)
	default:
		fmt.Println("default")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// receive
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("i3", i3)
		default:
			if n > 100 {
				break L
			}
		}
	}

}
