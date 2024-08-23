package main

import (
	f "fmt" // パッケージ名を変更可能
	// ."fmt" . であれば Println だけで使用可能
	"go-basic/scope"
)

func main() {
	f.Println(scope.ReturnMin())
	f.Println(scope.Max)

	// 小文字なのでprivate となり、別ファイルにて読み込みできない
	//fmt.Println(scope.min)

	f.Println(appName())
	Do("こんちは！")
}

func appName() string {
	const AppName = "go app name"
	var Version string = "1.0"
	return AppName + " " + Version
}

// Do https://go.dev/ref/spec#Blocks
func Do(a string) (b string) {
	b = a
	// {}の中に定義した値は {}のみ使用可能
	{
		b := "pppppp"
		f.Println(b) // pppppp
	}

	f.Println(b) //こんちは！
	return b
}
