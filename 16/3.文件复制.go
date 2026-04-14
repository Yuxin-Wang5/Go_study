package main

// 1.文件复制
io.Copy(dst Writer, src Reader) (written int64, err error)

//将src文件的内容复制到dst文件
read, _ := os.Open("go_study/file1.txt")
write, _ := os.Create("go_study/file3.txt") // 默认是可读可写，不存在就创建，清空文件
n, err := io.Copy(write, read)
fmt.Println(n, err)

// 2.目录操作
dir, _ := os.ReadDir("go_study")
for _, entry := range dir {
info, _ := entry.Info()
fmt.Println(entry.Name(), info.Size())    //文件名，文件大小，单位比特
}