package main

import (
	"fmt"
	"net"
	"sync"
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

	user := NewUser(conn)

	// 用户上线, 将用户加入到onlineMap中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	// 广播用户上线消息
	this.BroadCast(user, "已上线")

	select {}
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
