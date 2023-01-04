/**
 * @Author MierX
 * @Title 05多重赋值交换数据
 * @Date 2023-01-04 周三
 * @Time 20:52:01
 **/

package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println(a, b)

	// 多重赋值进行快速交换
	a, b = b, a
	fmt.Println(a, b)
}
