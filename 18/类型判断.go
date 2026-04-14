package main

//类型判断: 判断一个变量是否是结构体，切片，map

import (
	"fmt"
	"reflect"
)

func refType(obj any) {
	typeObj := reflect.TypeOf(obj)
	fmt.Println(typeObj, typeObj.Kind())
	// 去判断具体的类型
	switch typeObj.Kind() {
	case reflect.Slice:
		fmt.Println("切片")
	case reflect.Map:
		fmt.Println("map")
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.String:
		fmt.Println("字符串")
	}
}

func main() {
	refType(struct {
		Name string
	}{Name: "星"})
	name := "星"
	refType(name)
	refType([]string{"星"})
}
