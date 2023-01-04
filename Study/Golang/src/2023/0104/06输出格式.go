/**
 * @Author MierX
 * @Title 06输出格式
 * @Date 2023-01-04 周三
 * @Time 20:58:35
 **/

package main

import "fmt"

func main() {
	// PrintIn可以输出任何类型并独占一行
	fmt.Println(1)
	fmt.Println(2)

	// Print可以输出任何类型不独占一行
	fmt.Print(2)
	fmt.Print(3)

	// Printf只能支持输入字符串类型，可以通过设置占位符设置输出其他类型，不独占一行
	// %d 是一个占位符，表示输出一个整型数据
	// %f 是一个占位符，表示输出一个浮点型数据 默认保留六位小数，通过设置 %.nf 可以保留 n 位小数
	// \n 表示一个转义字符 换行
	// %s 是一个占位符，表示输出一个字符串类型
	a := 3
	b := 3.1415926
	fmt.Println()
	fmt.Printf("%d %.3f \n", a, b)
	fmt.Printf("%d %.3f \n", a, b)
	fmt.Printf("123")
}
