/*
 * @Author MierX
 * @Title 01数组.go
 * @Date 2023-02-02 周一
 * @Time 14:47:38
 */

package main

import "fmt"

func main0101() {
	// 数组定义
	// var 数组名 [元素个数]数据类型
	var arr [10]int

	// 通过下标找到具体元素 数组下标从0开始
	arr[0] = 123
	arr[1] = 110
	arr[2] = 234
	arr[5] = 999

	fmt.Println(arr)
}

func main0102() {
	// 在定义数组时 依次为数组元素赋值
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, i2 := range arr {
		fmt.Println(i, i2)
	}
}

func main0103() {
	// 在定义数组时 给指定下标赋值
	var arr = [10]int{1: 2, 3: 20, 8: 101}
	fmt.Println(arr)
}

func main0104() {
	// 在定义数组时 [...]可以根据元素个数赋值
	var arr = [...]int{1, 2, 3, 4}
	fmt.Println(arr)
}

func main() {
	// 数组常见问题

	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	arr[0] = 123
	fmt.Println(arr)
	// 超出数组下标的赋值会报错

	// 数组变量不能直接赋值

	// 两个数组的元素个数和类型相同，才能相互赋值

	// 打印数组类型
	fmt.Printf("%T\n", arr)

	// 打印数组地址
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[2])
}
