package main

import "fmt"

// https://go.dev/ref/spec#Interface_types

// StringFy sliceのなかにCarとPersonがある場合、それをfor文で回してメソッド発火したい時など便利
type StringFy interface {
	ToString() string
}

type Person struct {
	Name string
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v", p.Name)
}

type Car struct {
	Number string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v", c.Number)
}

func main() {
	vs := []StringFy{
		&Person{"John"},
		&Car{"12345"},
	}

	fmt.Println(vs) // [0xc000090020 0xc000090030]
	for _, v := range vs {
		fmt.Println(v.ToString())
		//Name=John
		//Number=12345
	}
}
