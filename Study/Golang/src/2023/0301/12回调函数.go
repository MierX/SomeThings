/*
 * @Author MierX
 * @Title 12回调函数
 * @Date 2023-03-08 周一
 * @Time 17:23:33
 */

package main

import "fmt"

type FuncType2 func(int int) int

// MyFunc20 回调函数，函数有一个参数是函数类型，这个函数就是回调函数
// 计算器 可以进行四则运算
// 多态 多种形态
func MyFunc20(a, b int, fTest FuncType2) (result int) {
	fmt.Println("MyFunc20")
	result = fTest(a, b)
	return
}

func MyFunc21(a, b int) int {
	return a + b
}

func main() {
	var fTest FuncType2
	fTest = MyFunc21
	fmt.Println(MyFunc20(10, 20, fTest))
}
