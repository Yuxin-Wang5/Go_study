package main

import "runtime"

// 1.一次性读取
byteData, _ := os.ReadFile("go_study/hello.txt")
fmt.Println(string(byteData))

// 2.获取当前go文件的路径：可以通过获取当前go文件的路径，然后用相对于当前go文件的路径去打开文件
// GetCurrentFilePath 获取当前文件路径
func GetCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// 3.分片读
file, _ := os.Open("go_study/hello.txt")
defer file.Close()
for {
	buf := make([]byte, 1)
	_, err := file.Read(buf)
	if err == io.EOF {
		break
	}
	fmt.Printf("%s", buf)
}

// 4.带缓冲读
//(1)按行读
file, _ := os.Open("go_study/hello.txt")
buf := bufio.NewReader(file)
for {
	line, _, err := buf.ReadLine()
	fmt.Println(string(line))
	if err != nil {
		break
	}
}

//(2)指定分割符
file, _ := os.Open("go_study/hello.txt")
scanner := bufio.NewScanner(file)
scanner.Split(bufio.ScanWords) // 按照单词读
//scanner.Split(bufio.ScanLines) // 按照行读
//scanner.Split(bufio.ScanRunes) // 按照中文字符读
//scanner.Split(bufio.ScanBytes) // 按照字节读读，中文会乱码
for scanner.Scan() {
	fmt.Println(scanner.Text())
}