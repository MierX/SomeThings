/*
 * @Author MierX
 * @Title 13数组做函数参数
 * @Date 2023-03-13 周一
 * @Time 16:23:55
 */

package main

import "fmt"

// 数组做函数参数，是值传递，实参数组的每个元素给形参数组拷贝一份
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify:a = ", a)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a) //数组传递过去
	fmt.Println("main:a = ", a)
}
