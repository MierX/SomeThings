/*
 * @Author MierX
 * @Title 07切片定义和使用.go
 * @Date 2023-02-03 周一
 * @Time 14:37:33
 */

package main

import "fmt"

func main0701() {
	// 数组定义：
	// var 数组名 [元素个数]数据类型
	// 切片定义：
	// var 切片名 []数据类型

	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Printf("%T\n", array)

	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)

	slice[0] = 100
	// 在使用切片时不能下标越界
	// slice[10] = 99

	// 利用append函数给切片添加下标
	slice = append(slice, 6, 7)
	fmt.Println(slice)
}

func main() {
	var slice []int
	fmt.Println(slice)
	fmt.Println(len(slice))

	// 在定义切片时 可以指定切片的长度
	var slice1 = make([]int, 10)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
}
