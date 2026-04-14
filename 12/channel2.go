package main

import (
	"fmt"
	"sync"
	"time"
)

//在同步模式下，channel没有任何意义,需要在异步模式下使用channel，在协程函数里面写，在主线程里面接收数据

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan <- money

	wait.Done()
}

// 协程
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	// 现在的模式，就是购物接力
	//shopping("张三")
	//shopping("王五")
	//shopping("李四")
	wait.Add(3)
	// 主线程结束，协程函数跟着结束
	go pay("张三", 2, &wait)
	go pay("王五", 3, &wait)
	go pay("李四", 5, &wait)

	go func() {
		defer close(moneyChan)
		// 在协程函数里面等待上面三个协程函数结束
		wait.Wait()
	}()

	for {
		money, ok := <-moneyChan
		fmt.Println(money, ok)
		if !ok {
			break
		}
	}

	//time.Sleep(2 * time.Second)

	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyChan)
}
