package main

//中断程序:遇到错误直接停止程序
//这种一般是用于初始化，一旦初始化出现错误，程序继续走下去也意义不大了，还不如中断掉

import (
	"fmt"
	"os"
)

func init() {
	// 读取配置文件中，结果路径错了
	_, err := os.ReadFile("xxx")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	fmt.Println("啦啦啦")
}
