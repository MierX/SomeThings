/*
 * @Author MierX
 * @Title 02普通参数列表
 * @Date 2023-03-01 周一
 * @Time 16:13:51
 */

package main

import "fmt"

//定义函数时，在函数名后面()定义的参数叫形参
//参数传递，只能由实参传递给形参，不能反过来，单向传递

// MyFunc01 有参无返回值函数的定义 普通参数列表
func MyFunc01(a int) {
	fmt.Println(a)
}

func MyFunc02(a int, b int) {
	fmt.Println(a, b)
}

func MyFunc03(a, b int) {
	fmt.Println(a, b)
}

func MyFunc04(a, b int, c string, d float64) {
	fmt.Println(a, b, c, d)
}

func main() {
	//有参无返回值函数调用：函数名（所需参数）
	//调用函数传递的参数叫实参
	MyFunc01(111)
	MyFunc02(111, 222)
	MyFunc03(111, 222)
	MyFunc04(111, 222, "abc", 3.14)
}
