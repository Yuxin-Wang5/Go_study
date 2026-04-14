package main

//线程安全下的map:
//如果在一个协程函数下，读写map就会引发一个错误
//concurrent map read and map write
//这个就是map的线程安全错误

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var mp = map[string]string{}

func reader() {
	for {
		fmt.Println(mp["time"])
	}
	wait.Done()
}
func writer() {
	for {
		mp["time"] = time.Now().Format("15:04:05")
	}
	wait.Done()
}

func main() {
	wait.Add(2)
	go reader()
	go writer()
	wait.Wait()

}

//我们不能在并发模式下读写map
//如果要这样做 ——> 1.加锁.go   2.sync.Map.go
