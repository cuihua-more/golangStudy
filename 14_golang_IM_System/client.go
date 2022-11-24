package main

import (
	"flag"
	"fmt"
	"net"
)

type Clinet struct {
	ServerIP   string
	ServerPort int

	Name string
	conn net.Conn
}

func NewClient(serverIp string, serverPort int) *Clinet {
	// 创建客户端对象
	client := &Clinet{
		ServerIP:   serverIp,
		ServerPort: serverPort,
	}

	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}

	client.conn = conn

	return client
}

var serverIp string
var serverPort int

// Usage: ./client -ip 127.0.0.1 -port 8888
func init() {
	// flag库：
	// StringVar方法表示从执行的程序选项中获取某一项内容，保存在传入的参数中
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器Port")
}
func main() {
	// 调用命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>> 连接服务器失败...")
		return
	}

	fmt.Println(">>>>>>>>>>> 连接服务器成功...")
	select {}
}
