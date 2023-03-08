/*
 * @Author MierX
 * @Title 09函数递归调用的流程
 * @Date 2023-03-08 周一
 * @Time 15:51:13
 */

package main

import "fmt"

func MyFunc15(a int) {
	//调用函数自身
	if a == 1 {
		fmt.Println("a = ", a)
		return
	}
	MyFunc15(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	MyFunc15(20)
}
