package main

import "fmt"

// 以年龄为例，输入的年龄在某一个区间，就输出对应的提示信息
// <=0     未出生
// 1-18    未成年
// 18-35   青年
// >=35    中年
//这是一个多选一的情况,我们有很多中方式来实现

func main() {
	fmt.Printf("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	//一、中断式，也叫卫语句,用return
	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}
	if age <= 35 {
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")

	//二、嵌套式
	//if age <= 18 {
	//    if age <= 0 {
	//      fmt.Println("未出生")
	//    } else {
	//      fmt.Println("未成年")
	//    }
	//  } else {
	//    if age <= 35 {
	//      fmt.Println("青年")
	//    } else {
	//      fmt.Println("中年")
	//    }
	//  }

	//三、多条件式 (无return)      &&:与      ||:或     ！：非
	//if age <= 0 {
	//    fmt.Println("未出生")
	//  }
	//  if age > 0 && age <= 18 {
	//    fmt.Println("未成年")
	//  }
	//  if age > 18 && age <= 35 {
	//    fmt.Println("青年")
	//  }
	//  if age > 35 {
	//    fmt.Println("中年")
	//  }

	//逻辑短路：
	// &&：第一个条件如果是false，后面的条件就不会去走了
	// ||：第一个条件如果是true，后面的条件就不会去走了

}
