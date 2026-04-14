package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Student struct {
	Name string `feng-orm:"name"`
	Age  int    `feng-orm:"age"`
}

type UserInfo struct {
	Id   int    `feng-orm:"id"`
	Name string `feng-orm:"name"`
	Age  int    `feng-orm:"age"`
}

func Find(obj any, query ...any) (sql string, err error) {
	// Find(Student, "name = ?", "fengfeng")
	// 希望能够生成 select name, age from  where  name = 'fengfeng'
	t := reflect.TypeOf(obj)
	//v := reflect.ValueOf(obj)
	// 首先得是结构体对吧
	if t.Kind() != reflect.Struct {
		err = errors.New("非结构体")
		return
	}
	// 拿全部字段

	// 拼接条件
	// 第二个参数，中的问号，就决定后面还能接多少参数
	var where string
	if len(query) > 0 {
		// 有第二个参数，校验第二个参数中的？个数，是不是和后面的个数一样
		q := query[0] // 理论上还要校验第二个参数的类型
		if strings.Count(q.(string), "?")+1 != len(query) {
			err = errors.New("参数个数不对")
			return
		}
		// 拼接where语句
		// 将？号带入后面的参数
		for _, a := range query[1:] {
			// 替换q
			// 这里要判断a的类型
			at := reflect.TypeOf(a)
			switch at.Kind() {
			case reflect.Int:
				q = strings.Replace(q.(string), "?", fmt.Sprintf("%d", a.(int)), 1)
			case reflect.String:
				q = strings.Replace(q.(string), "?", fmt.Sprintf("'%s'", a.(string)), 1)
			}
		}
		where += "where " + q.(string)
	}
	// 如果没有第二个参数，就是查全部

	// 拼接select
	// 拿所有字段，取feng-orm对应的值
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		f := field.Tag.Get("feng-orm")
		// 不考虑是空的情况
		columns = append(columns, f)
	}

	// 结构体的小写名字+s做表名
	name := strings.ToLower(t.Name()) + "s"

	// 拼接最后的sql
	sql = fmt.Sprintf("select %s from %s %s", strings.Join(columns, ","), name, where)
	return
}

func main() {
	sql, err := Find(Student{}, "name = ? and age = ?", "星", 3)
	fmt.Println(sql, err) // select name,age from students where name = '星' and age = 3
	sql, err = Find(UserInfo{}, "id = ?", 1)
	fmt.Println(sql, err) // select id,name,age from userinfos where id = 1
}

//对反射的一些建议:
//如果是写一下框架，偏底层工具类的操作，用反射确实不太好写，但是如果是在业务上，大量使用反射就不太合适了
//因为反射的性能没有正常代码高，会慢个一到两个数量级
//使用反射可读性也不太好，并且也不能在编译期间发生错误
