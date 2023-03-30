/*
 * @Author MierX
 * @Title 05http客户端
 * @Date 2023-03-30 周四
 * @Time 14:22:46
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	rs, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	defer rs.Body.Close()

	fmt.Println("Status =", rs.Status)
	fmt.Println("StatusCode =", rs.StatusCode)
	fmt.Println("Header =", rs.Header)
	//fmt.Println("Body =", rs.Body)
	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err1 := rs.Body.Read(buf)
		if n == 0 {
			fmt.Println("err1 =", err1)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("Body =", tmp)
}
