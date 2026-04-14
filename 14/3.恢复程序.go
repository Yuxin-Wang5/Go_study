package main

//恢复程序:我们可以在一个函数里面，使用一个defer，可以实现对panic的捕获,以至于出现错误不至于让程序直接崩溃
//这种一般也是框架层的异常处理所做的
//当然，这个用于捕获异常的defer的延迟函数可以在调用链路上的任何一个函数上
//一般用于在最上层函数，捕获所有异常

import (
	"fmt"
	"runtime/debug"
)

func read() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 捕获异常，打印错误信息
			// 打印错误的堆栈信息
			s := string(debug.Stack())
			fmt.Println(s)
		}
	}()
	var list = []int{2, 3}
	fmt.Println(list[2]) // 肯定会有一个panic
}

func main() {

	read()
}
