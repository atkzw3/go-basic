package main

import "fmt"

// go のデフォで下記が定義されている
//type error interface {
//	Error() string
//}

// MyError カスタム
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラー発生", ErrCode: 1}
}

func main() {
	er := RaiseError()
	fmt.Println(er)
}
