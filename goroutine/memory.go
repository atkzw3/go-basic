package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memConsumed() uint64 {
	// 使われていないメモリ開放
	runtime.GC()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m.Sys
}

// ゴルーチンのメリットはメモリが軽量であること
// この関数で計測する
// ゴルーチンは処理が終わるまでメモリ解放はされないので、発火させて計測する
func main() {

	var ch <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-ch
	}

	const numGoroutines = 100000

	wg.Add(numGoroutines)

	before := memConsumed()
	fmt.Println("before:", before)

	for i := 0; i < numGoroutines; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Println("after:", after)
	// Goroutine byte
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
