/*
 * @Author MierX
 * @Title 11浮点型
 * @Date 2023-02-27 周一
 * @Time 16:05:07
 */

package main

import "fmt"

func main() {
	//声明变量
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)

	//自动推导类型
	f2 := 3.14
	fmt.Println("f2 = ", f2)
}
