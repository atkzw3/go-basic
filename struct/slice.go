package main

import "fmt"

type SliceUser struct {
	Name string
}

type SliceUsers []*SliceUser

func main() {
	user1 := SliceUser{Name: "a"}
	user2 := SliceUser{Name: "b"}
	user3 := SliceUser{Name: "c"}

	users := SliceUsers{&user1, &user2, &user3}
	// 下記書き方でもOk
	//users := SliceUsers{}
	//users2 := append(users, &user1, &user2, &user3)
	fmt.Println(users) //[0xc000090020 0xc000090030 0xc000090040]
	for _, user := range users {
		fmt.Println(user)
		//&{a}
		//&{b}
		//&{c}
	}
}
