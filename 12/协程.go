package main

//Goroutine是Go运行时管理的轻量级线程
//在go中，开启一个协程是非常简单的

import (
	"fmt"
	"time"
)

func sing() {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
}

func main() {
	startTime := time.Now()
	go sing()
	go sing()
	go sing()
	go sing()
	time.Sleep(2 * time.Second)
	fmt.Println("唱歌结束", time.Since(startTime))
}

//如果把这个主线程中的延时去掉之后，会发现程序没有任何输出就结束了
//这是因为主线程结束协程自动结束，主线程不会等待协程的结束 ——> 只需要让主线程等待协程就可以了: WaitGroup
