package main

import (
	"fmt"
)

// https://go.dev/ref/spec#Switch_statements
// https://go.dev/ref/spec#Type_assertions

// 様々な型のものを渡してもswitch分で処理の分岐が可能
func anything(a interface{}) {
	fmt.Println("anything", a)
	switch v := a.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	default:
		fmt.Println("unknown")
	}
}

func main() {
	var x interface{} = 3
	// 型アサーションで、int型にする
	i := x.(int)
	fmt.Println(i)
	i1 := i + 3
	fmt.Println(i1)
	// interface + int なので計算できない
	//i2 := x + 3

	// x はintの値なので、エラーが出る
	//fl := x.(float64)
	//fmt.Println(fl) // panic: interface conversion: interface {} is int, not float64

	// エラーが出ないようにするには2つ目の返り値を用意する
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64) // 0 false

	if x == nil {
		fmt.Println("nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i)
	} else if s, isString := x.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("not a string")
	}

	// 上記同様のif条件と同様 ifより簡単にできる
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("not")
	}

	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "not aa")
	}

	anything("aaaa")
	anything(1)
}
