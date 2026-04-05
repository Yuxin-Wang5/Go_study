package main

import "fmt"

func main() {
	//自带换行
	fmt.Println("hello,xx", "aa")
	fmt.Println("bb", "cc")

	//格式化输出,不能自动换行
	fmt.Printf("%s 哇，你好美! ", "yaoyao")
	fmt.Printf("%s 哇，你好美! ", "琪琪")
	//加 “\n” 自动换行
	fmt.Printf("%s 哇，你好美! \n", "xingxing")
	fmt.Printf("%s 哇，你好美! \n", "heng")

	fmt.Printf("%d\n", 3)               // 整数
	fmt.Printf("%f\n", 3.1415926)       // 小数
	fmt.Printf("%.2f\n", 3.1415926)     // 保留2位小数
	fmt.Printf("%s\n", "哈哈哈")           // 字符串
	fmt.Printf("%v\n", "你好")            // 可以作为任何值的占位符输出
	fmt.Printf("%#v\n", "")             // 用go的语法格式输出，很适合打印空字符串
	fmt.Printf("%v\n", "")              // 会打印出空白
	fmt.Printf("%T %T\n", "qiqi", 1.25) // 打印类型

	//将格式化之后的内容赋值给一个变量
	name := fmt.Sprintf("%v", "qiqi")
	//或者：
	var f = fmt.Sprintf("%.2f\n", 3.1415926)
	fmt.Println(name)
	fmt.Println(f)

}
