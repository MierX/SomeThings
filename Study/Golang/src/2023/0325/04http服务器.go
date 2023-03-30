/*
 * @Author MierX
 * @Title 04http服务器
 * @Date 2023-03-30 周四
 * @Time 14:09:11
 */

package main

import (
	"fmt"
	"net/http"
)

// HandleConn w是给客户端回复数据的，req是读取客户端发送的数据
func HandleConn(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Method =", req.Method) //用户请求方式
	fmt.Println("URL =", req.URL)
	fmt.Println("Header =", req.Header)
	fmt.Println("Body =", req.Body)

	w.Write([]byte("hello~")) //给客户端回复数据
}

func main() {
	//注册处理函数，只要用户连接，自动调用指定的处理函数
	http.HandleFunc("/", HandleConn)

	//监听绑定
	http.ListenAndServe("127.0.0.1:8000", nil)
}
