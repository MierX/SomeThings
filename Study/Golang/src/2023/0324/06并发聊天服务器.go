/*
 * @Author MierX
 * @Title 06并发聊天服务器
 * @Date 2023-03-29 周一
 * @Time 16:30:47
 */

package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Client 客户端
type Client struct {
	C    chan string //用于发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

// 保存在线用户 cliAddr -> Client
var onlineMap map[string]Client

var message = make(chan string)

// Manager 转发消息
func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-message //没有消息前，此处阻塞

		//遍历Map，给每个在线成员发送消息
		for _, client := range onlineMap {
			client.C <- msg //将新消息写入用户的管道中
		}
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个用户结构体
	cli := Client{make(chan string), cliAddr, cliAddr}

	//把结构体加入到在线用户map中
	onlineMap[cliAddr] = cli

	//新开一个协程，用于给当前客户端发送信息
	go WriteMsgToClient(cli, conn)

	//用户上线广播
	message <- MakeMsg(cli, "login")

	//提示我是谁
	cli.C <- MakeMsg(cli, "I am here")

	//定义一个对方是否主动退出的管道
	isQuit := make(chan bool)

	//定义一个对方是否在发送数据
	hasData := make(chan bool)

	//新开一个协程，用于接收用户发送的消息并广播
	go func() {
		buf := make([]byte, 1024*2)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println(cli.Name, "conn.Read err =", err)
				return
			}

			//转发消息
			msg := string(buf[:n-1])

			if len(msg) == 3 && msg == "who" {
				//遍历在线用户列表并发送
				conn.Write([]byte("userList:\n"))
				for _, client := range onlineMap {
					conn.Write([]byte(client.Name + "\n"))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename successful\n"))
			} else {
				message <- MakeMsg(cli, msg)
			}

			hasData <- true
		}
	}()

	for {
		//通过select来检测channel的流通
		select {
		case <-isQuit:
			delete(onlineMap, cli.Addr)
			message <- MakeMsg(cli, "is logout")
			return
		case <-hasData:
		case <-time.After(30 * time.Second):
			//超时处理
			conn.Write([]byte("you is timeout"))
			delete(onlineMap, cli.Addr)
			message <- MakeMsg(cli, "timeout & leave")
			return
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		//给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}

	defer listener.Close()

	//新开一个协程，转发消息，只要有新消息来，给每个成员都发送此消息
	go Manager()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("listener.Accept err =", err1)
			continue
		}

		//处理用户连接
		go HandleConn(conn)
	}
}
