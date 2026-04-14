package main

//结构体反射
//读取json标签对应的值，如果没有就用属性的名称

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int `json:"age"`
}

func main() {
	s := Student{
		Name: "星",
		Age:  3,
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonField := field.Tag.Get("json")
		if jsonField == "" {
			// 说明json的tag是空的
			jsonField = field.Name
		}
		fmt.Printf("Name: %s, type: %s, json: %s, value: %v\n", field.Name, field.Type, jsonField, v.Field(i))
	}
}

//写的这个很简单，没有处理-和omitempty的情况
