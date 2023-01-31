/*
 * @Author MierX
 * @Title 07匿名函数.go
 * @Date 2023-01-31 周一
 * @Time 15:47:07
 */

package main

import "fmt"

func main() {
	a := 10
	b := 20

	// 在函数内部定义一个匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(a, b)

	v := func(a, b int) int {
		return a + b
	}(a, b)

	fmt.Println(v)
}
