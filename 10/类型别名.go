package main

import "fmt"

//类型别名:和自定义类型很像，但是有一些地方和自定义类型有很大差异
//1.不能绑定方法
//2.打印类型还是原始类型
//3.和原始类型比较，类型别名不用转换

type MyCode int        //无 = ,则为自定义类型
type MyAliasCode = int //有 = ，则为类型别名

// 自定义类型能绑定方法
func (m MyCode) Code() {

}

// 类型别名不能绑定方法
// func (n MyAliasCode)Code() {
//
// }
const myCode MyCode = 1
const myAliasCode MyAliasCode = 1

func main() {
	fmt.Printf("%v,%T\n", myCode, myCode)
	fmt.Printf("%v,%T\n", myAliasCode, myAliasCode)
	var age = 1
	if myCode == MyCode(age) {

	}
	if myAliasCode == age {

	}
}
