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
	flag int
}

func NewClient(serverIp string, serverPort int) *Clinet {
	// 创建客户端对象
	client := &Clinet{
		ServerIP:   serverIp,
		ServerPort: serverPort,
		flag:       999,
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

func (this *Clinet) menu() bool {
	var flag int

	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 更新用户名")
	fmt.Println("0. 退出")

	fmt.Scanln(&flag) // 从键盘获取

	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内的数字<<<<<")
		return false
	}
}

func (this *Clinet) Run() {
	for this.flag != 0 {
		for this.menu() != true {
		}

		// 根据不同的模式处理不同的业务
		switch this.flag {
		case 1:
			fmt.Println("公聊模式选择...")
			break
		case 2:
			fmt.Println("私聊模式选择...")
			break
		case 3:
			fmt.Println("更新用户名选择...")
			break
		}
	}
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

	client.Run()
}
