package main

// channel: 在协程里面产生了数据，传递给主线程,或者是传递给其他协程函数

import "fmt"

func main() {
	var c chan int // 声明一个传递整形的通道
	// 初始化通道
	c = make(chan int, 1) //  初始化一个 有一个缓冲位的通道
	c <- 1
	//c <- 2 // 会报错 deadlock
	fmt.Println(<-c) // 取值
	//fmt.Println(<-c) // 再取也会报错  deadlock

	c <- 2
	n, ok := <-c
	fmt.Println(n, ok)
	close(c) // 关闭协程
	c <- 3   // 关闭之后就不能再写或读了  send on closed channel
	fmt.Println(c)
}
