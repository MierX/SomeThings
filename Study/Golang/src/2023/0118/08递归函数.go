/*
 * @Author MierX
 * @Title 08递归函数.go
 * @Date 2023-01-31 周一
 * @Time 16:42:45
 */

package main

import "fmt"

// 递归函数
// 没有出口的递归函数是死递归 可能导致程序崩溃
func demo4(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	n--
	demo4(n)
}

func main() {
	demo4(10)
}
