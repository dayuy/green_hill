package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock sync.RWMutex

	// 消息广播的channel
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}

	return server
}

func (this *Server) Start() {
	// socket Listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port)) //"127.0.0.1:8888"
	if err != nil {
		fmt.Println("net.Listen err: ", err)
	}
	defer listener.Close()

	// 启动监听Message的goroutine
	go this.ListenMessage()

	// accept
	// 循环，表示当前main goroutine一直活着，保证了 sub goroutine（Handler()）活着
	for {
		conn, err := listener.Accept() // 连接成功，表示用户上线
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		// do handler
		go this.Handler(conn)
	}

	// close listen socket
}

func (this *Server) Handler(conn net.Conn) {
	// ...当前链接的业务
	fmt.Println("连接成功。。。")

	user := NewUser(conn, this)

	// 用户上线，将用户加入到onlineMap
	user.Online()
	//this.mapLock.Lock() // 操作onlineMap时，加锁
	//this.OnlineMap[user.Name] = user
	//this.mapLock.Unlock()

	// 广播当前用户上线的消息
	//this.BroadCast(user, "已上线")

	isLive := make(chan bool)
	// 接收客户端发送的消息并广播
	go this.FeedbackMess(user, conn, isLive)
	// 当前handler阻塞。来保证当前goroutine活着，子goroutine（FeedbackMess(user, conn)）正常工作
	//select {}

	// 长时间未活跃用户被踢出去
	for {
		select {
		case <-isLive:
			// 当前用户活跃中，重置定时器
			// break
			// 不做任何处理，会继续执行下面的case

		case <-time.After(time.Second * 30):
			user.SendMsg("您超时，被剔除了")

			// 销毁管道
			close(user.C)
			// 关闭连接
			conn.Close()  // conn.Read(buf) 0 触发Offline
			// 退出当前 Handler的goruntine, 来表示下线.
			return // runtime.Goexit()
		}
	}
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + ":" + msg
	// 向Server.Message管道发送信息，然后广播给所有用户
	// user.C <-sendMsg

	this.Message <- sendMsg
}

// 监听Message管道，有消息则分发给所有用户
func (this *Server) ListenMessage() {
	for {
		msg:=<-this.Message

		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 接收用户发送的消息 并广播
func (this *Server) FeedbackMess(user *User, conn net.Conn, isLive chan bool) {
	buf := make([]byte, 4096)
	for true {
		n, err := conn.Read(buf)
		if n == 0 {
			user.Offline()
			//this.BroadCast(user, "下线")
			return
		}

		if err!=nil && err != io.EOF {
			fmt.Println("Conn Read err:", err)
			return
		}

		// 提取用户的消息，去除'\n'
		msg := string(buf[:n-1])

		// 将得到的消息进行广播
		user.ReadMessage(msg)
		//this.BroadCast(user, msg)

		isLive <- true
	}
}
