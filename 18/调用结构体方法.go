package main

//调用结构体方法
//如果结构体有call这个名字的方法，就执行它

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (Student) See(name string) {
	fmt.Println("see name:", name)
}

func main() {
	s := Student{
		Name: "星",
		Age:  4,
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumMethod(); i++ {
		methodType := t.Method(i)
		fmt.Println(methodType.Name, methodType.Type)
		if methodType.Name != "See" {
			continue
		}
		methodValue := v.Method(i)
		methodValue.Call([]reflect.Value{
			reflect.ValueOf("星"), // 注意这里的类型
		})
	}
}
