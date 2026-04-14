package main

import (
	"fmt"
	"sync"
)

var num int
var wait sync.WaitGroup
var lock sync.Mutex

func add() {
	// 谁先抢到了这把锁，谁就把它锁上，一旦锁上，其他的线程就只能等着
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num++
	}
	lock.Unlock()
	wait.Done()
}
func reduce() {
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num--
	}
	lock.Unlock()
	wait.Done()
}

func main() {
	wait.Add(2)
	go add()
	go reduce()
	wait.Wait()
	fmt.Println(num)

}
