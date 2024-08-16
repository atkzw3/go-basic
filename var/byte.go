package main

import "fmt"

func main() {
	byteA := []byte("hello world")
	fmt.Println(byteA) // [104 101 108 108 111 32 119 111 114 108 100]

	// byteを文字列で出力
	sbyteA := string(byteA)
	fmt.Println(sbyteA) // hello world

	array := []byte{41, 2}
	fmt.Println(array) // [41 2]

	// rune 周り
	// Unicodeのコードポイントを表す
	// https://qiita.com/seihmd/items/4a878e7fa340d7963fee
	// https://zenn.dev/masaruxstudy/articles/52632501e4ca41

	s := "あ"

	for i := 0; i < len(s); i++ {
		b := s[i] // byte
		fmt.Println(b)
	}

	for i := 0; i < len(s); i++ {
		/* %xで16進数で出力する */
		fmt.Printf("% x", s[i]) // e3 81 82
	}

	for _, r := range s {
		// rune
		fmt.Println("\nrune")
		// Unicodeの番号を10進数に変換したもの
		fmt.Println(r) //12354
	}

}
