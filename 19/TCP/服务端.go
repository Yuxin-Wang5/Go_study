//TCP: 传输控制协议（TCP，Transmission Control Protocol）是一种面向连接的、可靠的、基于字节流的传输层通信协议

//如何保证连接可靠呢？（面试常考题）
//三次握手
//四次挥手

package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// 创建tcp的监听地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	// tcp监听
	listen, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		// 等待连接
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		// 获取客户端的地址
		fmt.Println(conn.RemoteAddr().String() + " 进来了")
		// 读取客户端传来的数据
		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			// 客户端退出
			if err == io.EOF {
				fmt.Println(conn.RemoteAddr().String() + " 出去了")
				break
			}
			fmt.Println(string(buf[0:n]))
		}

	}

}
