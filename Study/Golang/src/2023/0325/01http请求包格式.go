/*
 * @Author MierX
 * @Title 01http请求包格式
 * @Date 2023-03-30 周四
 * @Time 09:51:01
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户的连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err =", err1)
	}

	defer conn.Close()

	//接收客户端的数据
	buf := make([]byte, 4*1024)
	n, err2 := conn.Read(buf)
	if n == 0 {
		fmt.Println("conn.Read err =", err2)
		return
	}

	fmt.Printf("#\n%v#", string(buf[:n]))

	//#

	//请求行
	//GET / HTTP/1.1

	//请求头
	//Host: 127.0.0.1:8000
	//Connection: keep-alive
	//Cache-Control: max-age=0
	//sec-ch-ua: " Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"
	//sec-ch-ua-mobile: ?0
	//Upgrade-Insecure-Requests: 1
	//User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36
	//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
	//Sec-Fetch-Site: none
	//Sec-Fetch-Mode: navigate
	//Sec-Fetch-User: ?1
	//Sec-Fetch-Dest: document
	//Accept-Encoding: gzip, deflate, br
	//Accept-Language: zh-CN,zh;q=0.9

	//请求body
	//

	//#
}
