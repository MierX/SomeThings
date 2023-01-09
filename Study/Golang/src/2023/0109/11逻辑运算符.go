/*
 * @Author MierX
 * @Title 11逻辑运算符.go
 * @Date 2023-01-09 周一
 * @Time 15:23:37
 */

package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := a > b
	fmt.Println(c)
	// ! 是逻辑运算符 非 的符号，非真为假 非假为真
	fmt.Println(!c)

	// 单目运算符：只需要一个元素就能完成表达
	// 双目运算符：需要两个元素才能完成表达

	d := &a
	*d = 123
	fmt.Println(*d)
	fmt.Println(d)
	fmt.Println(a)
}
