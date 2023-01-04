/**
 * @Author MierX
 * @Title 09bool类型
 * @Date 2023-01-04 周三
 * @Time 22:11:10
 **/

package main

import "fmt"

func main() {
	// 布尔类型 值为 true 或者 false
	var a bool // 默认值为 false
	fmt.Println(a)
	a = true
	fmt.Println(a)

	// %T 是一个占位符 表示输出一个变量对应的数据类型
	fmt.Printf("%T", a)
}
