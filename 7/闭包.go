package main

import (
	"fmt"
	"time"
)

// &是取地址，
// *是解引用，去这个地址指向的值
// 设计一个函数，先传一个参数表示延时，后面再次传参数就是将参数求和
func awaitAdd(awaitSecond int) func(numberLise ...int) int {
	//time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(numberLise ...int) (sum int) {
		time.Sleep(time.Duration(awaitSecond) * time.Second) //2个地方延时后结果也不同
		for _, num := range numberLise {
			sum += num
		}
		return sum
	}
}

// 值传递
func Copy(name string) {
	fmt.Printf("infun %p\n", &name)
	name = "姚姚" //不影响name的值
}

// 传指针
func Set(name *string) {
	fmt.Printf("infun %p\n", name)
	*name = "星" //name为内存地址，不是字符串; *name为string类型，通过内存地址找值
}
func main() {
	//t1 := time.Now()
	//sum := awaitAdd(2)(1, 2, 3)
	//subTime := time.Since(t1)
	//fmt.Println(sum, t1, subTime)
	//add := awaitAdd(2)
	//t1 := time.Now()
	//sum := add(1, 2, 3)
	//subTime := time.Since(t1)
	//fmt.Println(sum, subTime)

	var name = "琪琪"
	Copy(name)
	fmt.Printf("%p\n", &name)
	Set(&name)
	fmt.Println(name)
}
