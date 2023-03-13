/*
 * @Author MierX
 * @Title 16切片的创建
 * @Date 2023-03-13 周一
 * @Time 17:09:30
 */

package main

import "fmt"

func main() {
	//切片和数组的区别
	//数组[]里面的长度是固定的一个常量，数组不能修改长度，len和cap永远都是固定的
	a := [5]int{}
	fmt.Printf("a:len = %d, cap = %d\n", len(a), cap(a))

	//切片，[]里面为空，或者为...，切片的长度或容量可以不固定
	s := []int{}
	fmt.Printf("s:len = %d, cap = %d\n", len(s), cap(s))
	s = append(s, 11)
	fmt.Printf("s:len = %d, cap = %d\n", len(s), cap(s))

	//自动推导类型，同时初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)

	//借助make函数，格式 make(切片类型，长度，容量)
	s2 := make([]int, 5, 10)
	fmt.Printf("s2:len = %d, cap = %d\n", len(s2), cap(s2))
}
