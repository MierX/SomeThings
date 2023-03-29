/*
 * @Author MierX
 * @Title 01TCP服务器
 * @Date 2023-03-29 周一
 * @Time 10:29:45
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000") //本地端口可以省略ip
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//进程结束后关闭监听
	defer listener.Close()

	//阻塞等待用户链接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//接收用户的请求
	buf := make([]byte, 1024) //1kb大小的缓冲区
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))

	defer conn.Close()
}
