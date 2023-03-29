/*
 * @Author MierX
 * @Title 接收文件
 * @Date 2023-03-29 周一
 * @Time 15:16:43
 */

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecFile(fileName string, conn net.Conn) {
	//新建文件
	file, err := os.Create("./Study/Golang/src/2023/0324/05文件传输/" + fileName)
	if err != nil {
		fmt.Println("os.Create err =", err)
		return
	}

	buf := make([]byte, 1024*4)

	//接收多少，写多少
	for {
		n, err1 := conn.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err =", err1)
			}
			return
		}

		if n == 0 {
			fmt.Println("n == 0，文件接受完毕")
		}

		//写入文件
		file.Write(buf[:n])
	}
}

func main() {
	//监听客户端请求
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err =", err)
	}

	defer listener.Close()

	//阻塞等待用户连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err =", err1)
		return
	}

	//接收文件名
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read err =", err2)
		return
	}

	fileName := string(buf[:n])

	//回复ok
	conn.Write([]byte("ok"))

	//接收文件内容
	RecFile(fileName, conn)
}
