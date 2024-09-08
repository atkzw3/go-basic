package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	var lock sync.Mutex

	const timer = 1 * time.Second

	freedByWorker := func() {
		defer wg.Done()

		count := 0
		begin := time.Now()

		// 全体をロック pointWorkerの方が回数が多い 効率が悪い
		for time.Since(begin) < timer {
			lock.Lock()
			time.Sleep(3 * time.Nanosecond)
			lock.Unlock()
			count++
		}

		fmt.Printf("greedByWorker : %v\n", count)
	}

	pointWorker := func() {
		defer wg.Done()

		count := 0
		begin := time.Now()

		// ロックの範囲は小さくなるが、ロック回数が増え取得するのが増える
		for time.Since(begin) < timer {
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()

			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()

			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()

			count++
		}

		fmt.Printf("pointWorker : %v\n", count)
	}

	wg.Add(2)
	go freedByWorker()
	go pointWorker()

	wg.Wait()

	/* 結論は、処理に合わせて適切にfreedByWorker　or pointWorker選択する*/
	/* 処理の時間計測などを比較して選択する*/
}
