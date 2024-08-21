package main

import "fmt"

type T struct {
	Employee Employee
}

type T2 struct {
	Employee
}

type Employee struct {
	Name string
}

func (e *Employee) SetName() {
	e.Name = "aiueo"
}

func main() {
	t := T{Employee: Employee{
		Name: "Jack",
	}}
	fmt.Println(t)               // {{Jack}}
	fmt.Println(t.Employee.Name) //Jack

	// T構造体に型だけ渡すこともできるが、その場合は取得方法は下記でも可能

	t2 := T2{Employee: Employee{
		Name: "Jee",
	}}
	fmt.Println(t2)               //{{Jee}}
	fmt.Println(t2.Name)          // Jee
	fmt.Println(t2.Employee.Name) //Jee

	t.Employee.SetName()
	fmt.Println(t.Employee.Name) //aiueo
}
