/*
 * @Author MierX
 * @Title 06gomaxprocs的使用
 * @Date 2023-03-28 周一
 * @Time 10:16:44
 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) //指定以单核运算
	fmt.Println("n = ", n)

	for {
		go fmt.Println(1)

		fmt.Println(0)
	}
}
