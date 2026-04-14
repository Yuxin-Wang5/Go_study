package main

import (
	"fmt"
	"reflect"
)

// 1.通过反射获取值
func refValue(obj any) {
	value := reflect.ValueOf(obj)
	fmt.Println(value, value.Type())
	switch value.Kind() {
	case reflect.Int:
		fmt.Println(value.Int())
	case reflect.Struct:
		fmt.Println(value.Interface())
	case reflect.String:
		fmt.Println(value.String())

	}
}

// 2.通过反射修改值
// 注意，如果需要通过反射修改值，必须要传指针，在反射中使用Elem取指针对应的值
func refSetValue(obj any) {
	value := reflect.ValueOf(obj)
	elem := value.Elem()
	// 专门取指针反射的值
	switch elem.Kind() {
	case reflect.String:
		elem.SetString("星穹列车")
	}
}

func main() {
	name := "星"
	refSetValue(&name)
	fmt.Println(name)
}
