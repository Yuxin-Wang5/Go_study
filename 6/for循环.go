package main

import "fmt"

func main() {
	//一、传统for循环
	//1+2+3+……100=？
	//var sum int
	//for i := 0; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//二、死循环：条件永远成立
	//for i := 0; true; i++ {
	//	fmt.Println(time.Now())
	//	time.Sleep(2 * time.Second)
	//}
	//死循环的简洁写法2：
	//for true {
	//	fmt.Println(time.Now())
	//	time.Sleep(2 * time.Second)
	//}
	//死循环的简洁写法2：
	//for {
	//	fmt.Println(time.Now())
	//	time.Sleep(1 * time.Second)
	//}

	//三、while模式：由于golang没有while循环，如果需要，则是由for循环稍微变化得来
	//while循环：先判断条件，再执行循环体
	//var sum int
	//var i int = 1
	//for i <= 100 {
	//	sum += i
	//	i++
	//}
	//fmt.Println(sum)

	//四、do-while模式:先执行一次循环体，再判断
	//var sum int
	//var i int = 1
	//for {
	//	sum += i
	//	i++
	//	if i == 101 {
	//		break
	//	}
	//}
	//fmt.Println(sum)

	//五、遍历切片:第一个参数是索引，第二个参数是项目
	//var list = []string{"七七", "瑶瑶", "星"}
	//法1：
	//for i := 0; i < len(list); i++ {
	//	fmt.Println(i, list[i])
	//}
	//法2：
	//for index, item := range list {
	//	fmt.Println(index, item)
	//}

	//六、遍历map
	//var userMap = map[string]string{"name": "瑶瑶"}
	//for key, value := range userMap {
	//	fmt.Println(key, value)
	//}

	//例如打印九九乘法表：
	// 1*1=1
	// 2*1=2  2*2=4
	// ……
	// 9*1=9  9*2=18  ……  9*9=81
	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	//	}
	//	fmt.Println()
	//}

	//七、break和continue：break用于跳出当前循环，continue用于跳过本轮循环
	//for i := 0; i <= 10; i++ {
	//	if i == 5 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

}
