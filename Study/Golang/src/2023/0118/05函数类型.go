/*
 * @Author MierX
 * @Title 05函数类型.go
 * @Date 2023-01-31 周一
 * @Time 14:52:04
 */

package main

import "fmt"

func demo1(a, b int) {
	fmt.Println(a + b)
}

// 定义函数类型 为已存在的数据类型起别名

type FUNCDEMO func(int, int)

func demo2(a, b int) {

}

func main() {
	// 函数的名字表示一个地址 函数在代码区的地址
	fmt.Println(demo1)

	f := demo1
	fmt.Println(&f)
	fmt.Printf("%T\n", f)
	f(1, 2)

	var f1 FUNCDEMO
	f1 = demo2
	fmt.Printf("%T\n", f1)
}
