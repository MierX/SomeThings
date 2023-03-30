/*
 * @Author MierX
 * @Title 03http响应包格式
 * @Date 2023-03-30 周四
 * @Time 13:41:06
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
		fmt.Println("err =", err)
		return
	}

	defer conn.Close()

	requestBuf := "GET /go HTTP/1.1\nHost: 127.0.0.1:8000\nConnection: keep-alive\nCache-Control: max-age=0\nsec-ch-ua: \" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"\nsec-ch-ua-mobile: ?0\nUpgrade-Insecure-Requests: 1\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\nSec-Fetch-Site: none\nSec-Fetch-Mode: navigate\nSec-Fetch-User: ?1\nSec-Fetch-Dest: document\nAccept-Encoding: gzip, deflate, br\nAccept-Language: zh-CN,zh;q=0.9\n"

	//先发请求包，服务器才会回应响应包
	conn.Write([]byte(requestBuf))

	//接收服务器回复的响应包
	buf := make([]byte, 4*1024)
	n, err1 := conn.Read(buf)
	if n == 0 {
		fmt.Println("err1 =", err1)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))
}
