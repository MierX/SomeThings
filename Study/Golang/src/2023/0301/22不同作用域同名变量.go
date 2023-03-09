/*
 * @Author MierX
 * @Title 22不同作用域同名变量
 * @Date 2023-03-09 周一
 * @Time 17:23:52
 */

package main

import "fmt"

var a byte //全局变量

func main() {
	var a int //局部变量
	//不同作用域，允许定义同名变量
	//使用变量的原则，就近原则
	fmt.Printf("%T\n", a)

	{
		var a float32
		fmt.Printf("2:%T\n", a)
	}
	MyFunc27()
}

func MyFunc27() {
	fmt.Printf("3:%T\n", a)
}
