package main

import "fmt"

// https://go.dev/ref/spec#Struct_types

type User struct {
	Name string
	Age  int
	// 複数のフィールド定義も可能
	x, y int
}

func main() {
	var u User
	fmt.Println(u) //{ 0 0 0}
	u.Name = "aiu"
	fmt.Println(u) // {aiu 0 0 0}
	u.x = 123
	fmt.Println(u) // {aiu 0 123 0}

	u2 := User{}
	fmt.Println(u2) //{ 0 0 0}
	u2.Name = "aiu"
	fmt.Println(u2) //{aiu 0 0 0}

	// 初期値も可能
	u3 := User{Name: "aiueo"}
	fmt.Println(u3) //{eee 0 0 0}

	// 初期値はフィールド名指定しなくても可能
	// この場合は全ての値は必須となる
	u4 := User{"oooo", 123, 22, 33}
	fmt.Println(u4) //{oooo 123 22 33}

	// ポインタ型となる
	u5 := new(User)
	fmt.Println(u5) //&{ 0 0 0}
	// ポインタ型となる
	// 使う場合はこちらが多い。構造体に渡す時はポインタ型が多い
	u6 := &User{}
	fmt.Println(u6) //&{ 0 0 0}

	updateUser(u6)
	fmt.Println(u6) // &{updateUser 0 0 0}
	r := updateUser2(u3)
	fmt.Println(u3) //{aiueo 0 0 0} ポインタ型で渡していないので更新されていない！！
	fmt.Println(r)  //{updateUser2 0 0 0} 関数からの値は更新される
}

func updateUser(u *User) {
	fmt.Println("updateUser")
	u.Name = "updateUser"
}

func updateUser2(u User) User {
	fmt.Println("updateUser2")
	u.Name = "updateUser2"
	return u
}
