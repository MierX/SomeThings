/*
 * @Author MierX
 * @Title 20局部变量
 * @Date 2023-03-09 周一
 * @Time 16:54:37
 */

package main

import "fmt"

func MyFunc26() {
	a := 10
	fmt.Println("a = ", a)
}

func main() {
	//定义在{}里面的变量就是局部变量，只能在{}里面有效
	//执行到定义变量那句话，才开始分配空间，离开作用域就被释放
	//作用域，变量起作用的范围

	//a = 111
	{
		i := 10
		fmt.Println("i = ", i)
	}
}
