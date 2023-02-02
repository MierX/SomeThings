/*
 * @Author MierX
 * @Title 02十只小猪称体重.go
 * @Date 2023-02-02 周一
 * @Time 16:00:13
 */

package main

import "fmt"

func main() {
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)

	// 求最重
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(max)

	// 数组的数据交换
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	fmt.Println(arr)
}
