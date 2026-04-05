package main

import "fmt"

//Go语言中的map(映射、字典)是一种内置的数据结构，它是一个无序的 key-value对的集合
//map的key必须是基本数据类型，value可以是任意类型
//注意:map使用之前一定要初始化！！！！！

func main() {
	var userMap map[int]string = map[int]string{
		1: "七七",
		2: "瑶瑶",
		6: "",
	}
	fmt.Println(userMap)
	fmt.Println(userMap[1])
	fmt.Println(userMap[2])
	//若要求输出不存在的元素，也不会报错，而是会直接输出此类型的默认值，eg: string的默认值为空字符串
	fmt.Println(userMap[4])
	//格式化输出会更明显
	fmt.Printf("%#v\n", userMap[5])

	//区分是map中的string类型的 6：“” 和不存在的空字符串
	fmt.Printf("%#v\n", userMap[6])
	value := userMap[6]     //string类型的 6：“”
	value, ok := userMap[6] //value指的是string类型的 6：“”    / ok指的是布尔值
	fmt.Println(value, ok)

	//变值
	userMap[1] = "行秋"
	fmt.Println(userMap)

	//删除元素
	delete(userMap, 6)
	fmt.Println(userMap)

}
