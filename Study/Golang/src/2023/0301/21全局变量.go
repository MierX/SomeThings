/*
 * @Author MierX
 * @Title 21全局变量
 * @Date 2023-03-09 周一
 * @Time 17:18:10
 */

package main

import "fmt"

// 定义在函数外部的变量是全局变量
// 全局变量在任何地方都能使用
var a int

func main() {
	a = 10
	fmt.Println(a)
}
