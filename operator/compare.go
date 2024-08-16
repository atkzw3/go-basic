package main

import "fmt"

// https://go.dev/ref/spec#Operators

func main() {
	is := 1 == 1
	fmt.Println(is)
	fmt.Printf("%T\n", is)

	notIs := 1 == 2
	fmt.Println(notIs)
	fmt.Printf("%T\n", notIs)

	compare1 := 4 <= 8
	compare2 := 4 >= 8
	compare3 := 4 < 8
	compare4 := 4 > 8
	fmt.Println(compare1)
	fmt.Println(compare2)
	fmt.Println(compare3)
	fmt.Println(compare4)

	true := true == true
	fmt.Println(true)

	true2 := false != true
	fmt.Println(true2)

}
