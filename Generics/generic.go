package main

import "fmt"

func PrintSliceInt(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func PrintSliceString(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// PrintSlice intとString　が同様処理を行っているケースを1つにまとめる場合
// https://recursionist.io/learn/languages/go/oop/generics
func PrintSlice[T int | string](a []T) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	PrintSliceInt([]int{1, 2, 3})
	PrintSliceString([]string{"A", "B", "C"})

	// int
	PrintSlice([]int{1, 2, 3})
	// string
	PrintSlice([]string{"A", "B", "C"})
}
