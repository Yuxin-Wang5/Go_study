package main

import (
	"fmt"
	"sync"
)

var num int
var wait sync.WaitGroup

func add() {
	for i := 0; i < 1000000; i++ {
		num++
	}
	wait.Done()
}
func reduce() {
	for i := 0; i < 1000000; i++ {
		num--
	}
	wait.Done()
}

func main() {
	wait.Add(2)
	go add()
	go reduce()
	wait.Wait()
	fmt.Println(num)

}

//什么是线程安全？
//现在有两个协程，同时触发，一个协程对一个全局变量进行100完成++操作，另一个对全局变量—的操作
//那么，两个协程结束，最后的值应该是0才对
//但是会发现，这个输出的结果完全无法预测
//这是为什么呢？
//根本原因是CPU的调度方法为抢占式执行，随机调度
//那能不能通过给操作加锁来解决这个问题呢?
//可以 ——> 同步锁.go
