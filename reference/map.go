package main

import "fmt"

// https://go.dev/ref/spec#Map_types

func main() {
	// keyは文字列 値は数値
	var m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m)

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2)

	m3 := map[int]string{1: "foo", 2: "bar"}
	fmt.Println(m3)

	// 空マップ作成
	m4 := make(map[int]string)
	fmt.Println(m4) // map[]

	m4[1] = "hello"
	m4[2] = "world"
	fmt.Println(m4) // map[1:hello 2:world]

	fmt.Println(m4[1]) // hello

	//　登録されていないkeyは空文字
	fmt.Println(m4[100]) // 空文字

	m5 := make(map[string]int)
	fmt.Println(m5)
	m5["aaaa"] = 123

	// 値がintの場合 keyがない場合は 0となる
	fmt.Println(m5["bbbb"])

	// 取得時のエラーハンドリング
	// 2つめの引数を用意する
	s, ok := m5["aaaa"]
	fmt.Println(s, ok) // 123 true

	s1, ok1 := m5["bbbb"]
	fmt.Println(s1, ok1) // 0 false

	if !ok1 {
		fmt.Println("エラーだよ")
	}

	m5["bbbb"] = 5555
	fmt.Println(m5["bbbb"])

	// 要素の削除
	delete(m5, "bbbb")
	fmt.Println(m5["bbbb"])
}
