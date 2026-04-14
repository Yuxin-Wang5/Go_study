package main

//修改结构体中某些值
//例如，结构体tag中有big的标签，就将值大写

import (
	"fmt"
	"reflect"
	"strings"
)

type Student struct {
	Name1 string `big:"-"`
	Name2 string
}

func main() {
	s := Student{
		Name1: "星",
		Name2: "穹",
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(&s).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		bigField := field.Tag.Get("big")
		// 判断类型是不是字符串
		if field.Type.Kind() != reflect.String {
			continue
		}
		if bigField == "" {
			continue
		}
		// 修改值
		valueFiled := v.Field(i)
		valueFiled.SetString(strings.ToTitle(valueFiled.String()))
	}
	fmt.Println(s)
}
