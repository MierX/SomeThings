/*
 * @Author MierX
 * @Title 05数组指针
 * @Date 2023-02-06 周一
 * @Time 15:59:42
 */

package main

import "fmt"

func main0501() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 定义指针 指向数组 数组指针
	p := &arr
	fmt.Printf("%T\n", p)

	var p1 *[10]int
	p1 = &arr
	// 通过指针间接操作数组
	for i := 0; i < len(p1); i++ {
		fmt.Println(p1[i])
	}

	fmt.Println(p1)
	p1[0] = 101
	fmt.Println(p1)
}

func main0502() {
	arr := [5]int{0, 1, 2, 3, 4}
	// 指针变量和要存储的数据类型相同
	// var p *[10]int = &arr

	// p1 和 p2 在内存中指向相同的地址 但是p1和p2的类型不同 一个是指向整个数组 一个是指向数组的元素
	p1 := &arr
	p2 := &arr[0]
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", p2)
}

func main0503() {
	arr := [5]int{0, 1, 2, 3, 4}
	p := &arr
	fmt.Println(arr)
	// 数组指针作为函数参数，是引用传递
	test2(p)
	fmt.Println(arr)
}

func test2(p *[5]int) {
	p[2] = 123
}

func BubbleSort(arr *[10]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := [10]int{0, 9, 4, 3, 1, 2, 5, 7, 6, 8}
	p := &arr
	fmt.Println(arr)
	BubbleSort(p)
	fmt.Println(arr)
}
