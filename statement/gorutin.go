package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("sub")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// sub()だけだと、無限ループのまま
	// goをつけることでmainの無限ループも発火する
	go sub()

	for {
		fmt.Println("main")
		time.Sleep(200 * time.Millisecond)
	}
}
