package main

import "fmt"

type Code int

func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "成功"
	case ServiceErrCode:
		return "服务错误"
	case NetworkErrCode:
		return "网络错误"
	}
	return ""
}

func (c Code) ok() (code Code, msg string) {
	return c, c.GetMsg()
}

const (
	SuccessCode    Code = 0
	ServiceErrCode Code = 1001 //服务错误
	NetworkErrCode Code = 1002 //网络错误
)

func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return ServiceErrCode.ok()
	}
	if name == "2" {
		return NetworkErrCode.ok()
	}
	return SuccessCode.ok()
}
func main() {
	fmt.Println(SuccessCode.GetMsg())
	var i int
	fmt.Println(int(SuccessCode) == i) // 必须要转成原始类型才能判断
}
