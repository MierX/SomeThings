/*
 * @Author MierX
 * @Title 13匿名函数和闭包
 * @Date 2023-03-09 周一
 * @Time 15:03:59
 */

package main

import "fmt"

func main() {
	a := 10
	str := "mike"

	//匿名函数，没有函数名字，函数定义，没有调用
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	f1()

	//给一个函数类型起别名
	type FuncType2 func() //函数没有参数没有返回值
	//声明变量
	var f2 FuncType2
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func() {
		fmt.Println(a)
	}()

	//定义带参数的匿名函数，同时调用
	func(a int) {
		fmt.Println(a)
	}(20)

	//定义带参数有返回值的匿名函数，同时调用
	b := func(a, b int) int {
		return a + b
	}(13, 23)
	fmt.Println(b)
}
