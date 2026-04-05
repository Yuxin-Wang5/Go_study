package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"名字"`
	Age      int    `json:"年龄,omitempty"`
	Password string `json:"-"`
}

func main() {
	user := User{Name: "七七", Age: 10, Password: "123456"}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))

}
