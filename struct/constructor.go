package main

import "fmt"

type User2 struct {
	Name string
}

// NewUser2 construct
func NewUser2(name string) *User2 {
	return &User2{Name: name}
}

func main() {
	u := NewUser2("world")
	fmt.Println(u) //&{world}
}
