/*
 * @Author MierX
 * @Title 02测试服务器
 * @Date 2023-03-30 周四
 * @Time 13:35:19
 */

package main

import (
	"fmt"
	"net/http"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello world")
}

func main() {
	http.HandleFunc("/go", myHandle)

	//在指定的地址进行监听，开启一个HTTP
	http.ListenAndServe(":8000", nil)
}
