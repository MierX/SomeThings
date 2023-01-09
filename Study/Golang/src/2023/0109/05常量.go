/**
 * @Author MierX
 * @Title 05常量
 * @Date 2023-01-08 周日
 * @Time 14:05:59
 **/

package main

import "fmt"

func main() {
	// 常量的定义 一般用大写字母命名
	const a int = 10
	b := a + 10
	fmt.Println(a)
	fmt.Println(b)

	// 常量不允许二次赋值
	// a = 100

	// 系统不能允许打印常量地址

	// 常量计算
	const PI float64 = 3.1415926
	var r float64
	fmt.Scan(&r)
	s := PI * r * r
	l := 2 * PI * r
	fmt.Printf("面积：%.2f\n", s)
	fmt.Printf("周长：%.2f\n", l)

	// 字面常量
	fmt.Println("hello world")
}
