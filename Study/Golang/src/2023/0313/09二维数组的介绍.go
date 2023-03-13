/*
 * @Author MierX
 * @Title 09二维数组的介绍
 * @Date 2023-03-13 周一
 * @Time 15:55:51
 */

package main

import "fmt"

func main() {
	//有多少个[]就是多少维
	//有多少个[]就是多少次循环
	var a [3][4]int
	k := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			k++
			a[i][j] = k
		}
	}
	fmt.Println("a = ", a)

	b := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println("b = ", b)
}
