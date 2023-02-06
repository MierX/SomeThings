/*
 * @Author MierX
 * @Title 03指针变量作为函数参数
 * @Date 2023-02-06 周一
 * @Time 15:49:44
 */

package main

import "fmt"

func main() {
	a := 10.9
	// 所有的指针类型都存储的是一个无符号十六进制整型数据
	p := &a
	fmt.Printf("%T\n", p)

	b := 10
	c := 20
	swap(&b, &c)
	fmt.Println(b, c)
}

func swap(a, b *int) {
	// *指针变量 是取值操作 解引用运算符
	// &变量 是取地址操作 引用运算符
	*a, *b = *b, *a
}
