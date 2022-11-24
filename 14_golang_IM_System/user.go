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
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式: to|name|msg
		spilStr := strings.Split(msg, "|")
		if len(spilStr) != 3 {
			this.SendMsg("消息格式不正确，使用 \"to|Zhang3|Hello\"\n")
			return
		}

		// 1. 获取对方用户名
		remoteName := spilStr[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确，使用 \"to|Zhang3|Hello\"\n")
			return
		}

		// 2. 获取已上线的User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg(remoteName + ": 不存在\n")
			return
		}

		// 3. 获取消息内容
		content := spilStr[2]
		if content == "" {
			this.SendMsg("无消息内容，请重发\n")
			return
		}
		remoteUser.SendMsg(this.Name + "say to you: " + content + "\n")
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
