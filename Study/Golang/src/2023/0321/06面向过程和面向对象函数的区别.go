/*
 * @Author MierX
 * @Title 06面向过程和面向对象函数的区别
 * @Date 2023-03-24 周一
 * @Time 15:16:41
 */

package main

import "fmt"

//实现两数相加

// Add01 面向过程
func Add01(a, b int) int {
	return a + b
}

// 面向对象，方法：给某个类型绑定一个函数
type long int

// Add02 tmp叫接收者，接收者就是传递的一个参数
func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = Add01(1, 1) //普通函数调用方式
	fmt.Println("result = ", result)

	//定义一个变量
	var a long
	//调用方法格式：变量名.函数(所需参数)
	rs := a.Add02(2)
	fmt.Println("rs =", rs)

	//面向对象只是换了一种表现形式
}
