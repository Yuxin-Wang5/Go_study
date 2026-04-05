package main

import "fmt"

type UserInfo struct {
	Name string "name"
}

func (this *UserInfo) SetName(name string) {
	this.Name = name
	fmt.Printf("this: %p\n", this)
}

func main() {
	user := UserInfo{
		Name: "七七",
	}
	fmt.Printf("user: %p\n", &user)
	user.SetName("瑶瑶")
	fmt.Println(user.Name)
}
