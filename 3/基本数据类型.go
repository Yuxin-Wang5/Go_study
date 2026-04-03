package main

import (
	"fmt"
	"math"
)

// 1. 默认的数字定义类型是int类型
// 2. 带个u就是无符号，只能存正整数
// 3. 后面的数字就是2进制的位数
// 4. uint8还有一个别名 byte， 一个字节=8个bit位
// 5. int类型的大小取决于所使用的平台
// 例如 uint8，那就是8个二进制位，都用来存储数据，那最小就是0，最大就是2的八次方-1=255
// 那 int8，因为要拿一位存符号，使用实际只有七位可用，所以最小的就是负2的七次方=-128，最大的就是2的七次方-1=127

func main() {
	//一、整数型
	var age = 12
	var u8 uint8 = 255
	//uint8
	// 00000000 = 0
	// 11111111 = 255 = 2^8-1

	//int8
	// 0 0000000 = 0
	// 0 1111111 = 127 = 2^7-1
	// 1 0000001 = -1(原码）
	// 1 0000000 = -128

	//第五点的测试: 64位操作系统，那么试一下 int 是不是就是 int64 的最大上限 ———— 2^63-1 = 9223372036854775807
	fmt.Printf("%.0f\n", math.Pow(2, 63))
	var n1 int = 9223372036854775807
	fmt.Println(n1)
	var n2 int = 9223372036854775808 // 看它报不报错
	fmt.Println(n2)

	//二、浮点型
	//Go语言支持两种浮点型数：float32 和 float64
	//float32 的浮点数的最大范围约为3.4e38，可以使用常量定义：math.MaxFloat32
	//float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64
	//如果没有显式声明，则默认是float64

	//三、字符型 ：注意是字符，不是字符串
	//比较重要的两个类型是 byte（单字节字符）和 rune（多字节字符）
	//在 Go 中，字符的本质是一个整数，直接输出时，是该字符对应的 UTF-8 编码的码值
	//可以直接给某个变量赋一个数字，然后按格式化输出时 %c ，会输出该数字对应的 unicode 字符
	//字符类型是可以进行运算的，相当于一个整数，因为它都对应有 Unicode 码
	//单字节
	var a byte = 'a' // ascii里面的字符 (ascii码其实就是uint8码，范围是0-255)   //用单引号
	fmt.Printf("%c %d\n", a, a)
	var a1 uint8 = 97
	fmt.Printf("%c %d\n", a1, a1)

	//多字节
	var z rune = '中'
	fmt.Printf("%c %d\n", z, z) // z = 200013: unicode码

	//四、字符串类型: 和字符不一样的是，字符的赋值是单引号，字符串的赋值是双引号
	var s string = "qiaqia"
	fmt.Println(s)

	//五、多行字符串：反引号（反引号中会原样输出，转义字符无用）
	fmt.Println(`hello \n
world`)
	fmt.Println(`hello 
go`)

	//六、转义字符
	//一些常用的转义字符
	fmt.Println("七七\t瑶瑶")     // 制表符
	fmt.Println("七七\n瑶瑶")     // 回车
	fmt.Println("\"七七\"瑶瑶")   // 双引号
	fmt.Println("七七\r瑶瑶")     // 回到行首
	fmt.Println("C:\\七七\\瑶瑶") // 反斜杠

	//七、布尔类型
	//布尔型数据只有 true（真）和 false（假）两个值
	//布尔类型变量的默认值为 false
	//Go 语言中不允许将整型强制转换为布尔型
	//布尔型无法参与数值运算，也无法与其他类型进行转换

	//八、零值问题
	//如果给一个基本数据类型只声明不赋值，那么这个变量的值就是对应类型的零值，例如int就是0，bool就是false，字符串就是""
	var c1 int
	var c2 float64
	var c3 string
	var c4 bool

	fmt.Printf("%#v\n", c1)
	fmt.Printf("%#v\n", c2)
	fmt.Printf("%#v\n", c3)
	fmt.Printf("%#v\n", c4)

}
