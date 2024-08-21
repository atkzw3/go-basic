package main

import "fmt"

type UserMap struct {
	Name string
}

func main() {
	m := map[int]UserMap{
		1: {Name: "A"},
		2: {Name: "B"},
	}
	fmt.Println(m) //map[1:{A} 2:{B}]

	m2 := map[UserMap]string{
		{Name: "A"}: "a",
		{Name: "B"}: "b",
	}
	fmt.Println(m2) // map[{A}:a {B}:b]
}
