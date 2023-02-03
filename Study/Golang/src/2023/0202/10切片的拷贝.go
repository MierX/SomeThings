/*
 * @Author MierX
 * @Title 10切片的拷贝.go
 * @Date 2023-02-03 周一
 * @Time 16:00:53
 */

package main

import "fmt"

func main() {
	var slice = []int{0, 1, 2, 3, 4}

	s := make([]int, 5)
	// copy函数用于为前一个参数复制后一个参数的值，且不共用内存地址
	copy(s, slice)
	fmt.Println(s) // [0 1 2 3 4]

	// copy出来的新切片发生变化，并不会影响原切片
	fmt.Println(slice) // [0 1 2 3 4]
	s[1] = 999
	fmt.Println(s)     // [0 999 2 3 4]
	fmt.Println(slice) // [0 1 2 3 4]

	s1 := make([]int, 10)
	copy(s1, slice)
	fmt.Println(s1) // [0 1 2 3 4 0 0 0 0 0]

	s2 := make([]int, 1)
	copy(s2, slice)
	fmt.Println(s2) // [0]
}
