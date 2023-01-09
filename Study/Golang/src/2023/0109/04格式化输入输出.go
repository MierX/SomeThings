/**
 * @Author MierX
 * @Title 04格式化输入输出
 * @Date 2023-01-08 周日
 * @Time 12:56:12
 **/

package main

import "fmt"

func main() {
	fmt.Printf("35%%\n")

	a := 10
	// %b 是一个占位符 打印数据的十进制
	fmt.Printf("%b\n", a)

	// %o 是一个占位符 打印数据的八进制
	fmt.Printf("%o\n", a)

	// %x 是一个占位符 打印数据的十六进制
	fmt.Printf("%x\n", a)

	// 十进制数据
	var b int = 10

	// 八进制数据
	var c int = 010

	// 十六进制数据
	var d int = 0xa

	// 二进制数据 不能在go语言中直接表示
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// 单引号括起来的只能存储一个字符
	e := 'a'
	fmt.Printf("%c\n", e)

	// %p 是一个占位符 表示打印一个变量对应的内存地址，是以无符号十六进制表示
	fmt.Printf("%p\n", &e)
}
