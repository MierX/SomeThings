/**
 * @Author MierX
 * @Title 04多层赋值和匿名变量
 * @Date 2023-01-04 周三
 * @Time 20:36:53
 **/

package main

import "fmt"

func main() {
	// 多重赋值
	a, b, c := 10, 20, 30
	fmt.Println(a, b, c)

	// 在一个作用域范围内变量定义变量名不能重复
	// var a int = 10
	// fmt.Println(a)

	// 匿名变量
	// _ 表示匿名变量，不接收数据
	_, c, d := 111, 222, 333
	// fmt.Println(_, c, d)
	fmt.Println(c, d)
}
