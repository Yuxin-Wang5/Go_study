package main

import "fmt"

// 定义函数
func sayHello() {
	fmt.Println("hello world")
}

func param1(id string) {
	fmt.Println(id)
}
func param2(id string, userName string) {
	fmt.Println(id, userName)
}

// 类型一样时，可以合并
func param3(id, userName string) {
	fmt.Println(id, userName)
}
func add(numberList ...int) {
	var sum int
	for _, i2 := range numberList {
		sum += i2
	}
	fmt.Println(sum)
}

// //没有返回值
func r1() {
	return
}

// 单返回值
func r2() bool {
	return false //也可以不写
	//或者
	var ok bool
	return ok
}

// 多返回值
//func r3(string, bool) {
//	if 1 > 2 {
//		return "", false
//	}
//	if 1 < 2 {
//		return "", false
//	}
//	return "", true
//}

// 命名返回值
func r4(val string, ok bool) {
	if 1 > 2 {
		val = "12"
		return
	}
	return
}

func main() {
	//sayHello()
	//param1("123")

	//add(1, 2)
	//add(1, 2, 3)

	//匿名函数：在函数里创建函数
	//var getName = func() string {
	//	return "爱爱"
	//}
	//var setName = func(name string) {
	//	fmt.Println(name)
	//	return
	//}
	//fmt.Println(getName())
	//setName("皮皮")

	//高阶函数：根据用户输入的不同，执行不同的操作
	fmt.Println("请输入要执行的操作：")
	fmt.Println("1.登录")
	fmt.Println("2.注册")
	fmt.Println("3.个人中心")
	var index int
	fmt.Scan(&index)

	var funMap = map[int]func(){
		1: login,
		2: register,
		3: userCenter,
	}
	fun, ok := funMap[index]
	if ok {
		fun()
	}
	//switch index {
	//case 1:
	//	login()
	//case 2:
	//	register()
	//case 3:
	//	userCenter()
	//}

}
func login() {
	fmt.Println("登录")
}
func register() {
	fmt.Println("注册")
}
func userCenter() {
	fmt.Println("个人中心")

}
