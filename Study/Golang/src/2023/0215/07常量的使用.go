/*
 * @Author MierX
 * @Title 07常量的使用
 * @Date 2023-02-27 周一
 * @Time 15:37:17
 */

package main

import "fmt"

func main() {
	//变量：程序运行期间，可以改变的量，变量声明需要var
	//常量：程序运行期间，不可以改变的量，常量声明需要const

	const a int = 10
	//a = 20 //err，常量不允许修改
	fmt.Println("a = ", a)
}
