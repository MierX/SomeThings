/*
 * @Author MierX
 * @Title 02TCP客户端
 * @Date 2023-03-29 周一
 * @Time 11:27:33
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer conn.Close()

	//发送数据
	conn.Write([]byte("are u ok?"))
}
