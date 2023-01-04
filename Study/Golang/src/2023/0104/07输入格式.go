/**
 * @Author MierX
 * @Title 07输入格式
 * @Date 2023-01-04 周三
 * @Time 21:33:01
 **/

package main

import "fmt"

func main() {
	var a int

	// 通过键盘为变量赋值
	// & 是一个运算符 取地址运算符
	fmt.Scan(&a)

	// 打印变量内存地址 0xc00000e098 是一个十六进制整数型数据
	//fmt.Println(&a)

	fmt.Println(a)

	// 为多个变量输入，用空格或回车表示一个变量输入的结束
	var b, c int
	fmt.Scan(&b, &c)
	fmt.Println(b, c)

	var d int
	var f string
	// 在接受字符串时要使用空格作为分割
	fmt.Scanf("%d%s", &d, &f)
	fmt.Println(d, f)

	// 通过键盘输入学生三门成绩计算总成绩和平均数
	var g, h, i int
	fmt.Scan(&g, &h, &i)
	sum := g + h + i
	fmt.Println("总成绩：", sum)
	fmt.Println("平均成绩：", sum/3)
}
