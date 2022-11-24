package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 当前服务器在线用户的表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 添加用户列表方法
func (this *Server) AddUser(user *User) {
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
}

// 删除用户列表
func (this *Server) DelUser(user *User) {
	this.mapLock.Lock()
	delete(this.OnlineMap, user.Name)
	this.mapLock.Unlock()
}

// 监听广播Message Channel的goroutine，一旦有消息，就发送给全部在线的User
func (this *Server) ListernMessage() {
	for {
		msg := <-this.Message

		// 将消息发送给所有在线的User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.c <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播上线消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// 当前业务
	// fmt.Println("链接建立成功")

	user := NewUser(conn, this)

	user.Online()

	// 用作消息活跃
	isAlive := make(chan bool)

	// 从客户端读取数据
	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// 如果读对端返回0，表示链接已断开
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}

			// 定义一个字符串，内容是将字节流byte类型转化为string，同时使用slice不要最后一个，也就是"\n"
			msg := string(buf[:n-1])
			user.DoMessage(msg)
			isAlive <- true
		}
	}()

	// 将此函数阻塞
	for {
		select {
		case <-isAlive:
			// 什么也不做，会接着执行下面的case，但不执行case满足后的情况
		case <-time.After(time.Second * 10): // 重新开始定时10s
			// 如果进入到这，表示10秒超时
			// 关闭User，回收资源
			user.SendMsg("你被踢了")
			close(user.c)
			conn.Close() // 关闭socket

			return // runtime.Goexit
		}
	}
}

// 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	defer listener.Close()

	// 启动监听广播Message的goroutine
	go this.ListernMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error: ", err)
		}
		go this.Handler(conn)
	}

}
