/*
 * @Author MierX
 * @Title 06二维数组的定义使用.go
 * @Date 2023-02-03 周一
 * @Time 11:33:34
 */

package main

import "fmt"

func main0601() {
	// 定义二维数组
	var arr [2][3]int
	fmt.Println(len(arr))
	fmt.Println(len(arr[0]))

	arr[0][1] = 123
	arr[1][2] = 234
	fmt.Println(arr)

	// 外层控制行 内层控制列
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}

	for _, ints := range arr {
		for _, v := range ints {
			fmt.Println(v)
		}
	}
}

func main() {
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr)

	var arr1 = [2][3]int{1: {1, 2, 3}}
	fmt.Println(arr1)
}
