package main

import "fmt"

func main() {
	fmt.Println("请输入你的名字0：")
	//或者————光标可以直接在冒号后面
	fmt.Print("请输入你的名字1：")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)

	fmt.Print("请输入你的年龄：")
	var age int
	fmt.Scan(&age)
	fmt.Println(age)
}
