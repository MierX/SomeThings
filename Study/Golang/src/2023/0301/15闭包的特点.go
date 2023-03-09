/*
 * @Author MierX
 * @Title 15闭包的特点
 * @Date 2023-03-09 周一
 * @Time 15:13:48
 */

package main

import "fmt"

// MyFunc24 函数的返回值是一个匿名函数，返回一个函数类型
func MyFunc24() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func MyFunc23() int {
	var x int //没有赋值，初始化为0
	x++
	return x * x
} //函数调用完毕，x被销毁

func main() {
	fmt.Println(MyFunc23()) //1
	fmt.Println(MyFunc23()) //1
	fmt.Println(MyFunc23()) //1

	//返回值是一个匿名函数，通过f来调用返回来的匿名函数，f调用闭包函数
	//闭包函数不关心这谢谢捕获了的变量和常量是否已经超出作用域，只要有闭包还在使用它，这些变量还会存在
	f := MyFunc24()
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9

	f1 := MyFunc24()
	fmt.Println(f1()) //1
}
