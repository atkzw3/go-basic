package main

import (
	"fmt"
	"sync"
)

// https://go-tour-jp.appspot.com/concurrency/9
// https://qiita.com/atsutama/items/4f8e64ca7de9ce4b9064

func main() {
	var i int
	var wg sync.WaitGroup

	go func() {
		i++
	}()

	if i == 0 {
		fmt.Println("i is 0")
	} else {
		fmt.Println("i is 1")
	}

	// ↓ 今回でいうクリティカルセクションは i のこと。
	// 上記はgoroutineは順番が保証されていないので 0 がでたり1が出たりする
	// クリティカルセクションを考慮する
	// コンピュータ上において、単一の計算資源（リソース）に対して、複数の処理が同時期に実行されると、破綻をきたす部分を指す。クリティカルセクションにおいては、排他制御を行なうなどしてアトミック性を確保する必要がある。
	// https://ja.wikipedia.org/wiki/%E3%82%AF%E3%83%AA%E3%83%86%E3%82%A3%E3%82%AB%E3%83%AB%E3%82%BB%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3

	/* 解決方法 */

	// デメリット　ロックしてしまうので、パフォーマンスが落ちる

	var mA sync.Mutex
	var i2 int

	wg.Add(1)
	go func() {
		defer wg.Done()
		mA.Lock()
		i2++
		mA.Unlock()
	}()

	wg.Wait()

	mA.Lock()
	if i2 == 0 {
		fmt.Println("i2 is 0")
	} else {
		fmt.Println("i2 is 1")
	}
	mA.Unlock()
}
