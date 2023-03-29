/*
 * @Author MierX
 * @Title TCP服务器
 * @Date 2023-03-29 周一
 * @Time 11:34:29
 */

package main

import (
	"fmt"
	"net"
	"strings"
)

// HandleConn 处理用户请求
func HandleConn(conn net.Conn) {
	//函数调用完毕关闭连接
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connect successful")

	buf := make([]byte, 2*1024)

	for {
		//读取用户数据
		n, errMsg := conn.Read(buf)
		if errMsg != nil {
			fmt.Println("errMsg = ", errMsg)
			return
		}
		str := string(buf[:n-1])

		if strings.ToLower(str) == "exit" {
			fmt.Println(addr, "exit")
			return
		}

		fmt.Printf("[%s]：%s\n", addr, str)

		//把数据转换为大写再发送给用户
		conn.Write([]byte("收到信息了"))
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	//接收多个用户的请求
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("err1 = ", err1)
			continue
		}

		//处理用户请求，每个用户新建一个协程
		go HandleConn(conn)
	}
}
