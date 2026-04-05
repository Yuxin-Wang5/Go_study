package main

import (
	"fmt"
	"sort"
)

func main() {
	var nameList [3]string = [3]string{"琪琪", "七七", "瑶瑶"}
	var nameList1 [3]int = [3]int{1, 2, 3}
	fmt.Println(nameList)
	fmt.Println(nameList1)
	//         "琪琪"  "七七"  "瑶瑶"
	//索引（正向） 0       1      2      ————正向从0开始
	//索引（负向） -3      -2     -1     ————负向从-1开始
	fmt.Println(nameList[0])
	fmt.Println(nameList[1])
	fmt.Println(nameList[2]) // go语言不支持负向索引，但Java支持
	fmt.Println(len(nameList))
	fmt.Println(nameList[len(nameList)-1])

	// go语言里的数组长度被限制死了，可以修改元素值，但不能添加元素
	nameList[0] = "不死途"
	fmt.Println(nameList)

	//切片（Slice）相较于数组更灵活，因为在声明切片后其长度是可变的
	var List []string
	List = append(List, "绯英")
	List = append(List, "银狼")
	fmt.Println(List[0])
	fmt.Println(List[1])

	var List1 []string
	//List = append(List, "绯英")
	//List = append(List, "银狼")
	fmt.Println(List1 == nil)

	var List2 []string = []string{} //法2
	//List = append(List, "绯英")
	//List = append(List, "银狼")
	fmt.Println(List2 == nil)

	var List3 = []string{} //法3
	List4 := []string{}    //法4
	//法5
	var List5 []string = []string{}
	List5 = make([]string, 0)
	//List = append(List, "绯英")
	//List = append(List, "银狼")
	fmt.Println(List3)
	fmt.Println(List4)
	fmt.Println(List5)

	//可以通过make函数创建知道长度，指定容量的切片: make([]type, length, capacity)
	agelist := make([]int, 3)
	fmt.Println(agelist)

	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices)
	fmt.Println(array[0:2])
	fmt.Println(array[1:2])

	//切片排序
	var list = []int{4, 5, 3, 2, 7}
	fmt.Println("排序前:", list)
	sort.Ints(list) //默认升序
	fmt.Println("升序:", list)
	sort.Sort(sort.Reverse(sort.IntSlice(list))) //若想降序，则反转
	fmt.Println("降序:", list)
}
