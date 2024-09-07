package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
	name  string
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("aaa")

	wg.Add(2)

	var a value = value{name: "a"}
	var b value = value{name: "b"}

	fmt.Println("あああ")
	go printSum(&a, &b, &wg)
	go printSum(&b, &a, &wg)

	go wg.Wait()
}

func printSum(v1 *value, v2 *value, wg *sync.WaitGroup) {

	defer wg.Done()
	//v1.mu.Lock()

	fmt.Printf("%v がロックを取得しました\n", v1.name)
	//defer v1.mu.Unlock()

	// 適当な処理を行っている想定
	time.Sleep(2 * time.Second)

	//v2.mu.Lock()
	fmt.Printf("%v がロックを取得しました\n", v2.name)
	// v2.mu.Unlock()

	fmt.Println(v1.value, v2.value)
}
