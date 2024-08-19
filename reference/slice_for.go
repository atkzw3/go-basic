package main

import "fmt"

func main() {
	sl := []string{"a", "b", "c", "d"}
	fmt.Println(sl)

	for _, v := range sl {
		fmt.Println(v)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
