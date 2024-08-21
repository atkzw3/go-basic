package main

import "fmt"

type Admin struct {
	Name string
	Age  int
}

// SayName (u User)は レシーバーを定義している
// イメージとして、PHPのAdminクラスに定義しているメソッドのような使い方ができる main関数参照
func (u Admin) SayName() {
	fmt.Println(u.Name)
}

func (u Admin) SetName(name string) {
	u.Name = name
}

func (u *Admin) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := Admin{Name: "user1"}
	user1.SayName()

	user1.SetName("rrr")
	user1.SayName() // user1のまま  値渡ししているので更新できていない

	// 更新したい場合は、関数の型 をポイント型に変更する
	// 出来るだけレシーバーはポインタ型にしておくこと
	user1.SetName2("hi!!")
	user1.SayName() // hi!!

	// ポインタ型で定義することでも更新できるが、関数に定義するだけでできるので
	user2 := &Admin{Name: "user2"}
	user2.SayName() // user2
	user2.SetName2("world")
	user2.SayName() //world

}
