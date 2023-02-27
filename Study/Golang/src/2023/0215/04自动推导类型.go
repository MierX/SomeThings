/*
 * @Author MierX
 * @Title 04自动推导类型
 * @Date 2023-02-27 周一
 * @Time 15:15:16
 */

package main

import "fmt"

func main() {
	//赋值前，必须先声明变量
	var a int
	a = 10
	a = 20
	a = 30 // 赋值可以使用n次
	fmt.Println("a = ", a)

	b := 20
	fmt.Println("b = ", b)

	//b := 30 前面已经有变量b，不能再新建一个变量b
}
