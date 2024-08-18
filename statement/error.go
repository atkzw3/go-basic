package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "100"

	i, _ := strconv.Atoi(s)
	fmt.Println(i)

	var s1 string = "A"
	i, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("error:", err) // error: strconv.Atoi: parsing "A": invalid syntax
	}
	fmt.Println(i)
}
