/*
 * @Author MierX
 * @Title 03指针
 * @Date 2023-02-06 周一
 * @Time 15:08:01
 */

package main

import "fmt"

func main0301() {
	// 栈：先进后出
	// 队：先进先出

	a := 10
	fmt.Printf("%p\n", &a)

	// 定义指针变量p，存储变量a的内存地址
	p := &a
	fmt.Println(p)

	// 通过指针变量p改变变量a的值
	*p = 123
	fmt.Println(a)
}

func main0302() {
	// 声明指针变量p，默认值为0x0（nil）
	// 内存地址编号为0 0~255的空间为系统占用 不允许用户访问（读写）
	var p *int
	fmt.Println(p)
	fmt.Printf("%p\n", p)

	// new(数据类型)：开辟数据类型对应的内存空间，返回值为数据类型指针（内存地址）
	// gc 垃圾回收机制
	p = new(int)
	fmt.Println(p)
	fmt.Printf("%p\n", p)

	*p = 123
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Printf("%p\n", p)
}

func main() {
	// 野指针：指针变量指向了未知的空间
	// var p *int = *int(0xc042058088)
	// 指针变量必须有一个合理的指向
	// 在程序中允许出现空指针，不允许出现野指针
}
