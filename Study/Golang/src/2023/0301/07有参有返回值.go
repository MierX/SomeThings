/*
 * @Author MierX
 * @Title 07有参有返回值
 * @Date 2023-03-01 周一
 * @Time 16:52:18
 */

package main

import "fmt"

func MyFunc12(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}

func main() {
	fmt.Println(MyFunc12(12, 20))

	//通过匿名变量丢弃某个返回值
	a, _ := MyFunc12(12, 20)
	fmt.Println(a)
}
