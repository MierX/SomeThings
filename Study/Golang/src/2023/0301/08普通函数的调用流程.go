/*
 * @Author MierX
 * @Title 08普通函数的调用流程
 * @Date 2023-03-08 周一
 * @Time 15:36:42
 */

package main

import "fmt"

func MyFunc13() (a int) {
	a = 111

	//调用另外一个函数
	b := MyFunc14()
	fmt.Println("MyFunc14 b = ", b)

	fmt.Println("MyFunc13 a = ", a)
	return
}

func MyFunc14() (b int) {
	b = 222
	return
}

func main() {
	MyFunc13()
}
