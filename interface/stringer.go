package main

import "fmt"

// Stringer go に元々あるメソッド
// print.go
type Stringer interface {
	String() string
}

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B) // <<100, hello>>
	//return fmt.Sprintf("(%d, %q)", p.A, p.B)
}

func main() {
	p := &Point{A: 100, B: "hello"}
	fmt.Println(p) // &{100 hello}

}
