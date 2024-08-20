package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	for k := range m {
		fmt.Println(k)
	}
}
