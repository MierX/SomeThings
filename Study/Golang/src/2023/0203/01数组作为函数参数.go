/*
 * @Author MierX
 * @Title 01数组作为函数参数.go
 * @Date 2023-02-03 周一
 * @Time 16:18:49
 */

package main

import "fmt"

func main() {
	a := 10
	b := 20

	// 值传递，函数内部的形参发生改变，不会改变调用函数传入的实参
	swap(a, b) // 此时a,b属于传入的实参
	fmt.Println(a, b)

	arr := [10]int{9, 8, 4, 5, 3, 6, 7, 1, 2, 0}
	// 数组作为函数的参数，同样是值传递
	arr = BubbleSort(arr)
	fmt.Println(arr)
}

func swap(a, b int) { // 函数定义的参数，是形参
	a, b = b, a
}

func BubbleSort(arr [10]int) [10]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
