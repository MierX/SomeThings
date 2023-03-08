/*
 * @Author MierX
 * @Title 11函数类型
 * @Date 2023-03-08 周一
 * @Time 16:27:17
 */

package main

import "fmt"

func MyFunc18(a, b int) int {
	return a + b
}

func MyFunc19(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数类型起名

// FuncType FuncType代表它是一个函数类型
type FuncType func(int, int) int //没有函数名字，没有{}

func main() {
	var result int
	result = MyFunc18(1, 1) //传统调用方式
	fmt.Println(result)

	//声明一个函数类型的变量，变量名叫fTest
	var fTest FuncType
	fTest = MyFunc18
	result = fTest(10, 20)
	fmt.Println(result)
}
