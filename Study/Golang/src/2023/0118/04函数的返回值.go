/*
 * @Author MierX
 * @Title 04函数的返回值.go
 * @Date 2023-01-31 周一
 * @Time 14:15:18
 */

package main

import "fmt"

// func 函数名 (函数形参列表) 函数返回值类型 | (函数返回值1类型, 函数返回值2类型) {代码体}
func test4(a, b int) int {
	sum := a + b
	return sum
}

// 多个返回值
func test5(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	a := 10
	b := 20
	sum := test4(a, b)
	fmt.Println(sum)

	// 函数有多个返回值时接收变量要一一对应
	sum1, sub := test5(20, 30)
	fmt.Println(sum1, sub)
	fmt.Println(test5(50, 50))
}
