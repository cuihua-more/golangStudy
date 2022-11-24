package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	c    chan string
	conn net.Conn

	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		c:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动channel消息监听，并发送
	go user.ListenMessage()

	return user
}

// 用户上线业务
func (this *User) Online() {
	// 用户上线, 将用户加入到server.onlineMap中
	this.server.AddUser(this)

	// 广播用户上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户下线业务
func (this *User) Offline() {
	// 用户下线，从server.onlineMap删除
	this.server.DelUser(this)

	// 广播用户下线消息
	this.server.BroadCast(this, "已下线")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Name + "]" + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[0:7] == "rename|" {
		// 消息格式：rename|张三
		newName := strings.Split(msg, "|")[1] // 字符串分割

		//判断那么是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg(newName + ": 已经使用")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("以更新用户名:" + this.Name + "\n")
		}
	} else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前User channle的方法，一旦有消息，就发送给对端
func (this *User) ListenMessage() {
	for {
		msg := <-this.c

		this.conn.Write([]byte(msg + "\n"))
	}
}
