/**
 * @Author MierX
 * @Title 03自动类型推导
 * @Date 2023-01-04 周三
 * @Time 15:51:09
 **/

package main

import "fmt"

func main() {
	a := 10           // int 整型
	b := "hello word" // string 字符串型
	c := 123.456      // float64 浮点型
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	main004()

	// 不同的类型数据之间不能计算，如：a + c
}

func main004() {
	a := 10
	b := 20

	c := a
	a = b
	b = c

	fmt.Println(a, b)
}
