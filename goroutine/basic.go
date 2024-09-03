package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	// 遅延で終了させる
	defer wg.Done()
	fmt.Println("hello world")
}

// https://zenn.dev/kyoheiu/articles/b716964924298b
func main() {
	/*
		下記だとhello worldは出力されない
		main関数が先に終わってしまうため、時間的にgo sayHelloが出力されないため
		go sayHello()
	*/

	// 発火を起動するためには、sync.WaitGroupを使用する
	var wg sync.WaitGroup
	// goroutine前におく。goroutineのタスク登録するイメージ
	wg.Add(1)
	go sayHello(&wg)

	// goroutineが終わるまで待機
	wg.Wait()
}
