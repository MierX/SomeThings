/*
 * @Author MierX
 * @Title 01函数的定义和使用.go
 * @Date 2023-01-18 周一
 * @Time 15:44:42
 */

package main

import "fmt"

//	func 函数名(参数列表)(返回值列表){
//			代码体
//	}
//
// 在整个项目中，函数名是唯一的，不能重名
func add(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {
	a := 10
	b := 20

	// 函数调用 函数可以被多次调用
	add(a, b)
	add(1, 3)
}
