package main

import "syscall"

// 1.一次性写
err := os.WriteFile("go_study/file1.txt", []byte("这是内容"), os.ModePerm)
fmt.Println(err)

// 2.文件的打开方式
//(1)常见的一些打开模式:
// 如果文件不存在就创建
os.O_CREATE|os.O_WRONLY
// 追加写
os.O_APPEND|os.O_WRONLY
// 可读可写
os.O_RDWR

//(2)完整的模式
const (
	O_RDONLY int = syscall.O_RDONLY // 只读
	O_WRONLY int = syscall.O_WRONLY // 只写
	O_RDWR   int = syscall.O_RDWR   // 读写

	O_APPEND int = syscall.O_APPEND // 追加
	O_CREATE int = syscall.O_CREAT  // 如果不存在就创建
	O_EXCL   int = syscall.O_EXCL   // 文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 同步打开
	O_TRUNC  int = syscall.O_TRUNC  // 打开时清空文件
)

// 3.文件的权限
//主要用于linux系统，在windows下这个参数会被无视，代表文件的模式和权限位
//三个占位符:
//第一个是文件所有者所拥有的权限; 第二个是文件所在组对其拥有的权限; 第三个占位符是指其他人对文件拥有的权限
//根据 UNIX系统的权限模型，文件或目录的权限模式由三个数字表示，分别代表所有者（Owner）、组（Group）和其他用户(Other)的权限。每个数字由三个比特位组成，分别代表读、写和执行权限。因此，对于一个mode参数值，它的每一个数字都是一个八进制数字，代表三个比特位的权限组合
R: 读, Read的缩写, 八进制值为 4
W: 写, Write的缩写, 八进制值为 2
X: 执行, Execute的缩写, 八进制值为 1
//例如:
0444 表示三者均为只读的权限
0666 表示三者均为"读写"的权限
0777 表示三者均为读写执行的权限
0764 表示所有者有读写执行(7=4+2+1)的权限, 组有读写(6=4+2)的权限, 其他用户则为只读（4=4）