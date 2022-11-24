package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Clinent struct {
	ServerIP   string
	ServerPort int

	Name string
	conn net.Conn
	flag int
}

func NewClient(serverIp string, serverPort int) *Clinent {
	// 创建客户端对象
	client := &Clinent{
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

// 处理server端回应的消息，直接显示到标准输出
func (this *Clinent) DealResponse() {
	io.Copy(os.Stdout, this.conn) // io.copy方法会不断的调用this.conn总的read方法，读取处理，然后调用os.Stdout中的write方法，将数据写入
}

func (this *Clinent) menu() bool {
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

func (this *Clinent) UpdateName() bool {
	fmt.Println("请输入用户名：")
	fmt.Scanln(&this.Name)

	sendMsg := "rename|" + this.Name + "\n"

	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("更新用户名失败")
		return false
	} else {
		return true
	}
}

func (this *Clinent) Run() {
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
			// fmt.Println("更新用户名选择...")
			this.UpdateName()
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

	// 单独开启一个goroutine，
	go client.DealResponse()
	fmt.Println(">>>>>>>>>>> 连接服务器成功...")

	client.Run()
}
