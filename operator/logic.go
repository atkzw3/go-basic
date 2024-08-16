package main

import "fmt"

// https://go.dev/ref/spec#Operators
func main() {
	false := true && false == true
	fmt.Println(false)

	fmt.Println(!true)
}
