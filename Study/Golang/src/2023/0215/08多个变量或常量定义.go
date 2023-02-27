/*
 * @Author MierX
 * @Title 08多个变量或常量定义
 * @Date 2023-02-27 周一
 * @Time 15:45:49
 */

package main

import "fmt"

func main() {
	//不同类型变量的声明（定义）
	var a int
	var b float64
	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//多个不同类型变量声明快速方法
	var (
		c int
		d float64
	)
	c, d = 10, 3.14
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	//这样也可以自动推导类型
	const (
		e = 10
		f = 3.14
	)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
}
