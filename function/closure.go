package main

import "fmt"

func Later() func(string) (string, string) {
	var store string
	return func(next string) (string, string) {
		s := store
		store = next
		return s, next
	}
}

func main() {

	// クロージャーは実行時には値を前回の値で保持している性質がある
	f := Later()
	fmt.Println(f("aaaa")) // 空文字 aaaa
	fmt.Println(f("wwww")) // aaaa wwww
	fmt.Println(f("zzz"))  // wwww zzz
}
