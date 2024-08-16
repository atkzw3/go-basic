package main

import "fmt"

// https://go.dev/ref/spec#Operators
func main() {
	fmt.Println(1 + 2) // 3
	fmt.Println(2 - 2) // 0
	fmt.Println(2 * 2) // 4
	fmt.Println(2 / 2) // 1
	fmt.Println(3 % 2) // 1

	fmt.Println("AB" + "CD") // ABCD
	n := 0
	n += 4
	fmt.Println(n) // 4
	n++
	fmt.Println(n) // 5
	n--
	fmt.Println(n) //4

	s := "abcd"
	s += "e"
	fmt.Println(s) // abcde

}
