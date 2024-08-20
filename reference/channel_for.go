package main

import "fmt"

func main() {
	ch1 := make(chan int, 5)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	close(ch1)
	// closeしないままだとデータがない場合あるので、deadlockとなる
	// forで使用する場合はcloseを入れること
	for i := range ch1 {
		fmt.Println(i)
	}

}
