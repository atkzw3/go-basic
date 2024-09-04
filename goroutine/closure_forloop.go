package main

import (
	"fmt"
	"sync"
)

func main() {
	//var wg sync.WaitGroup
	//
	//say := "hello"
	//wg.Add(1)
	//
	////
	//go func() {
	//	defer wg.Done()
	//	say = "world"
	//}()
	//
	//wg.Wait()
	//// worldとなる (上書き)
	//fmt.Println(say)

	var wg2 sync.WaitGroup

	tasks := []string{"A", "B", "C"}
	for _, task := range tasks {
		wg2.Add(1)

		// 引数で渡さないと、発火までの時間でCCCと表示されるので注意
		go func(task string) {
			defer wg2.Done()
			fmt.Println(task)
		}(task)
	}

	wg2.Wait()
}
