package main

import "fmt"

func main() {
	var age int
	fmt.Printf("请输入你的年龄：")
	fmt.Scan(&age)

	switch {
	case age <= 0:

		fmt.Println("未出生")
	case age <= 18:
		fmt.Println("未成年")
		//fallthrough: 可让代码继续往下走
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("未成年")
	}

	fmt.Printf("请输入星期(数字)：")
	var week int
	fmt.Scan(&week)

	switch week {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
	case 4:
		fmt.Println("周四")
	case 5:
		fmt.Println("周五")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("错误")
	}

}
