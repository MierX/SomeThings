/*
 * @Author MierX
 * @Title 10bool类型
 * @Date 2023-02-27 周一
 * @Time 16:00:49
 */

package main

import "fmt"

func main() {
	//1、声明变量
	var a bool
	a = true
	fmt.Println("a = ", a)

	//2、自动推到类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)
}
