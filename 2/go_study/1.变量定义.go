package main

import (
	_ "2/go_study/version"
	"fmt"
	"strconv"
)

func hello() {
	fmt.Println("hello world")
	fmt.Println(age)
}

// 函数外（无论是什么函数，包括main）的变量是全局变量
// 函数外变量声明不能用短声明符号，必须用关键字声明，eg: age1 := 45 是错误的
var age string = strconv.Itoa(05)

var (
	c1 string = "55"
	c2 string = "66"
)

// 定义常量
const version string = "2.0.0"

// 常量名大写可跨包使用————version -> version.go
const Version1 string = "2.0.1"

// 定义在函数体（包括main函数）内的变量都是局部变量，定义了就必须使用；而全局变量可以只定义不使用
func main() {
	fmt.Println(version)
	//fmt.Println(version.yaoyao)

	//版本1：
	//先声明
	var name0 string
	//再赋值
	name0 = "123"
	fmt.Println(name0)

	//版本2：声明并赋值
	var name1 string = "45"
	fmt.Println(name1)

	//版本3：声明并赋值,省略类型
	var name2 = "67"
	fmt.Println(name2)

	//版本4：声明并赋值——短声明符号
	name3 := "89"
	fmt.Println(name3)

	hello()

	//定义多个变量
	//方法1
	var a1, a2 = 11, 22
	fmt.Println(a1, a2)

	//方法2
	var (
		b1 string = "33"
		b2 string = "44"
	)
	fmt.Println(b1, b2)

	fmt.Println(c1, c2)

}
