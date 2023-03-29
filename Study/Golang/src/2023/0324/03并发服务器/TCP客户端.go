/*
 * @Author MierX
 * @Title TCP客户端
 * @Date 2023-03-29 周一
 * @Time 12:00:42
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	//main调用完毕，关闭连接
	defer conn.Close()

	//从键盘输入内容，给服务器发送内容
	go func() {
		for {
			str := make([]byte, 1024)
			n, err2 := os.Stdin.Read(str)
			if err2 != nil {
				fmt.Println("os.Stdin.Read err = ", err2)
				return
			}

			//把输入的内容发送给服务器
			conn.Write(str[:n])
		}
	}()

	//接收服务器回复的数据
	//切片缓存
	buf := make([]byte, 1024)
	for {
		//接收服务器的请求
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("conn.Read err = ", err1)
			return
		}

		//打印接收到的内容
		fmt.Println(string(buf[:n]))
	}
}
