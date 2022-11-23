package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	c    chan string
	conn net.Conn
}

// 创建一个用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		c:    make(chan string),
		conn: conn,
	}

	// 启动channel消息监听，并发送
	go user.ListenMessage()

	return user
}

// 监听当前User channle的方法，一旦有消息，就发送给对端
func (this *User) ListenMessage() {
	for {
		msg := <-this.c

		this.conn.Write([]byte(msg + "\n"))
	}
}
