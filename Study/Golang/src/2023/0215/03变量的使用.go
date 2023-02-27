/*
 * @Author MierX
 * @Title 03变量的使用
 * @Date 2023-02-27 周一
 * @Time 14:45:29
 */

package main

import "fmt"

func main() {
	//变量，程序运行期间，可以改变的量

	//1、声明格式 var 变量名 类型，变量声明了，必须要使用
	//2、只是声明没有初始化的变量，默认值为0
	//3、同一个{}里，声明的变量名是唯一的
	var a int
	fmt.Println("a = ", a)

	//4、可以同时声明多个变量
	//var b, c int

	a = 10 //变量的赋值
	fmt.Println("a = ", a)

	//2、变量的初始化，声明变量时，同时赋值
	var b = 10 //初始化，声明变量时，同时赋值（一步到位）
	b = 20     //赋值，先声明，后赋值
	fmt.Println("b = ", b)

	//3、自动推导类型，必须初始化，通过舒适化的值确定类型
	c := 30
	fmt.Printf("c type is %T\n", c)
}
