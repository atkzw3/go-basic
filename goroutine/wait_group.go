package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Addはゴルーチン直前で書くこと！
	wg.Add(1)
	go func() {
		// 処理が終わったらaddした1を-1する
		// 必ずdeferするように！！
		defer wg.Done()
		fmt.Println("goroutine start")
		time.Sleep(time.Second)
		fmt.Println("goroutine end")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2 goroutine start")
		time.Sleep(time.Second)
		fmt.Println("2 goroutine end")
	}()

	var CPU int = runtime.NumCPU()

	fmt.Println("CPU:", CPU)

	// まとめて起動するのもOK
	wg.Add(CPU)
	for i := 0; i < CPU; i++ {
		go Hello(&wg)
	}

	// addしたものが0になるまで待つ
	wg.Wait()
}

func Hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}
