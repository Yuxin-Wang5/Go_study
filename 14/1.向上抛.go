package main

//go的异常处理可能是这门语言唯一的一个诟病了
//由于go语言没有捕获异常的机制，导致每调一个函数都要接一下这个函数的error

//常见的异常处理
//向上抛:将错误交给上一级处理,一般是用于框架层，有些错误框架层面不能擅做决定，将错误向上抛不失为一个好的办法

import (
	"errors"
	"fmt"
)

func Parent() error {
	err := method() // 遇到错误向上抛
	return err
}
func method() error {
	return errors.New("出错了")
}

func main() {
	fmt.Println(Parent())
}
