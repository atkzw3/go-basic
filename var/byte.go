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
}
