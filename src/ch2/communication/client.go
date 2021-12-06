package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp string
	ServerPort int
	conn net.Conn
	flagNum int
	Name string
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
		flagNum: 999,
	}

	// 在网络network上连接地址serverIP:Port。返回一个Conn接口
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error", err)
		return nil
	}

	client.conn = conn

	return client
}

func (client *Client) menu() bool {
	var flagNum int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flagNum) // 用户输入

	if flagNum >= 0 && flagNum <= 3 {
		client.flagNum = flagNum
		return true
	} else {
		fmt.Println(">>>>请输入合法范围内的数字<<<<")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>> 请输入名字：")
	fmt.Scanln(&client.Name) // 用户输入后，赋值给client.Name

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.write err:", err)
		return false
	}
	return true
}

func (client *Client) PublicChat() {
	var chatMsg string

	fmt.Println(">>>请输入聊天内容，exit退出.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		 if len(chatMsg) != 0 {
			 sendMsg := chatMsg + "\n"
			 _, err := client.conn.Write([]byte(sendMsg))
			 if err != nil {
				 fmt.Println("conn Write err: ", err)
				 break
			 }
		 }

		 chatMsg = ""
		 fmt.Println(">>>请输入聊天内容，exit退出。")
		 fmt.Scanln(&chatMsg)

		 // 如果输入为exit则退出此for循环
	}
}

func (client *Client) SelectUser() {
	sendMess := "who\n"
	_, err := client.conn.Write([]byte(sendMess))
	if err != nil {
		fmt.Println("conn Write err: ", err)
		return
	}
}

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUser()
	fmt.Println(">>>请输入聊天对象「用户名」，exit退出：")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>> 请输入消息内容，exit退出：")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err: ", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>请输入聊天内容，exit退出。")
			fmt.Scanln(&chatMsg)

			// 如果输入为exit则退出此for循环
		}

		client.SelectUser()
		fmt.Println(">>>>请输入聊天对象【用户名】，exit退出")
		fmt.Scanln(&remoteName)

		// 如果输入的内容为exit，则退出此for循环
	}
}

func (client *Client) DealResponse() {
	// 一旦 client.conn 有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)

	// 等同于
	//for {
	//	buf := make([]byte, 1024)
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}
}

func (client *Client) Run() {
	for client.flagNum != 0 {
		for client.menu() != true {
		}

		switch client.flagNum {
		case 1:
			fmt.Println("公聊模式选择...")
			client.PublicChat()
			break
		case 2:
			fmt.Println("私聊模式...")
			client.PrivateChat()
			break
		case 3:
			fmt.Println("更新用户名")
			client.UpdateName()
			break
		}
	}

	// == 0 就结束执行
}

var serverIp string
var serverPort int

// 通过命令行 获取参数IP、PORT
// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器地址")
	flag.IntVar(&serverPort, "port",8888, "设置服务器端口")
}

func main() {
	//client := NewClient("127.0.0.1", 8888)

	// 使用命令行参数
	flag.Parse()
	fmt.Println(serverIp, serverPort)

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>> 连接失败")
		return
	}
	fmt.Println(">>>> 连接成功...")

	go client.DealResponse()

	// 启动客户端业务
	//select {}
	client.Run()
}
