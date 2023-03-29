/*
 * @Author MierX
 * @Title 发送文件
 * @Date 2023-03-29 周一
 * @Time 15:03:46
 */

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err =", err)
		return
	}

	defer file.Close()

	//读文件内容，读多少发多少
	buf := make([]byte, 1024*4)
	for {
		n, err1 := file.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("file.Read err =", err1)
			}
			return
		}

		//发送内容
		conn.Write(buf[:n])
	}
}

func main() {
	//提示输入文件
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)

	//读取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err =", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", ":8000")
	if err1 != nil {
		fmt.Println("net.Dial err =", err1)
		return
	}

	defer conn.Close()

	//给服务器先发送文件名
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.Write err =", err2)
		return
	}

	//接收对方的回复，如果回复”ok“，说明可以发文件
	buf := make([]byte, 1024)

	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Read err =", err3)
		return
	}

	if string(buf[:n]) == "ok" {
		//发送文件内容
		SendFile(path, conn)
	}
}
