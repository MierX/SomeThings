/*
 * @Author MierX
 * @Title 02切片作为函数参数.go
 * @Date 2023-02-03 周一
 * @Time 16:32:29
 */

package main

import "fmt"

func main() {
	slice := []int{9, 8, 4, 5, 3, 6, 7, 1, 2, 0}

	fmt.Println(slice)        // [9 8 4 5 3 6 7 1 2 0]
	fmt.Printf("%p\n", slice) // 0xc00001a140

	// 切片本身是一个地址，所以在调用函数传入切片时，不再是值传递，而是引用传递
	// 函数内部对切片形参进行改变时，函数外部传入的切片实参也会被改变
	test(slice)
	fmt.Println(slice) // [0 1 2 3 4 5 6 7 8 9]

	// 如果在调用函数中对形参切片进行数据添加，不改变实参切片
	test1(slice)
	fmt.Println(slice) // [0 1 2 3 4 5 6 7 8 9]
}

// 切片作为函数参数
func test(slice []int) {
	//fmt.Println(slice) // [9 8 4 5 3 6 7 1 2 0]
	//fmt.Printf("%p\n", slice) // 0xc00001a140

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func test1(slice []int) {
	slice = append(slice, 10, 11)
}
