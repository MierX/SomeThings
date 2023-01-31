/*
 * @Author MierX
 * @Title 06函数的作用域.go
 * @Date 2023-01-31 周一
 * @Time 15:22:55
 */

package main

import "fmt"

// 在函数外部定义的变量 称为全局变量
// 作用域是项目中整个文件去使用
// 定义的全局变量不能和其他文件中的全局变量重名
var a int = 10

func main0601() {
	// 变量先定义 后使用 在函数内部变量名是唯一的
	// 在函数内部定义的变量 称为局部变量 局部变量的作用域在函数内部

	var i int = 10
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0 1 2 3 4
	}

	fmt.Println(i) // 10
}

func main() {
	// 变量先定义 后使用 在函数内部变量名是唯一的
	// 在函数内部定义的变量 称为局部变量 局部变量的作用域在函数内部

	var i int = 10
	for i = 0; i < 5; i++ {
		fmt.Println(i) // 0 1 2 3 4
	}

	fmt.Println(i) // 10

	// 全局变量可以和局部变量重名
	a := 123
	fmt.Println(a)

	// 修改全局变量的值
	a = 110
	demo3()
}

func demo3() {
	fmt.Println(a)
}
