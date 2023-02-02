/*
 * @Author MierX
 * @Title 03冒泡排序.go
 * @Date 2023-02-02 周一
 * @Time 17:18:46
 */

package main

import "fmt"

func main() {
	// 冒泡排序 升序
	count := 0
	var arr = [10]int{9, 1, 6, 5, 3, 2, 4, 8, 7, 10}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			count++
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	fmt.Println(count)
}
