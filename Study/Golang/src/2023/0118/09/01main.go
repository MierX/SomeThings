/*
 * @Author MierX
 * @Title 01main.go
 * @Date 2023-02-02 周一
 * @Time 14:01:32
 */

package main

import (
	"2023/0118/09/userinfo"
)

func main() {
	test(10, 20)
	demo()

	userinfo.UserCreate()
	userinfo.UserLogin()
}
