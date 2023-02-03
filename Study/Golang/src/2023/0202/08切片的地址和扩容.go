/*
 * @Author MierX
 * @Title 08切片的地址和扩容.go
 * @Date 2023-02-03 周一
 * @Time 14:52:59
 */

package main

import "fmt"

func main0801() {
	// 切片名本身就是一个引用地址
	// 创建的空切片 指向内存地址编号为0的空间
	var slice []int
	fmt.Printf("%p\n", slice)

	// 当使用append进行追加数据时 切片地址可能会发生改变
	slice = append(slice, 0, 1, 2)
	fmt.Printf("%p\n", slice)
	slice = append(slice, 3, 4, 5)
	fmt.Printf("%p\n", slice)

	var p *int
	fmt.Println(p)
	fmt.Printf("%p\n", p)
}

func main0802() {
	// make函数创建切片的参数：数据类型 长度 容量
	// 容量是为变量在内存地址中预留的空间
	var slice = make([]int, 10)
	fmt.Println(slice)
	fmt.Printf("%p\n", slice) // 0xc00001a140

	slice1 := make([]int, 10, 20)
	fmt.Println(slice1)
	fmt.Printf("%p\n", slice1) // 0xc0000220a0

	slice[0] = 100
	fmt.Println(slice)
	fmt.Printf("%p\n", slice) // 0xc00001a140

	slice1[0] = 100
	fmt.Println(slice1)
	fmt.Printf("%p\n", slice1) // 0xc0000220a0

	// 没有预设容量的切片长度发生变化后，内存地址可能会变化
	slice = append(slice, 1)
	fmt.Println(slice)
	fmt.Printf("%p\n", slice) // 0xc000022140

	// 有预设容量的切片长度发生变化后，长度未超过容量时，内存地址不会改变
	slice1 = append(slice1, 1)
	fmt.Println(slice1)
	fmt.Printf("%p\n", slice1) // 0xc0000220a0

	// 有预设容量的切片长度发生变化后，长度超过容量时，内存地址可能会改变
	slice1 = append(slice1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17)
	fmt.Println(slice1)
	fmt.Printf("%p\n", slice1) // 0xc00007e000
}

func main() {
	// 在使用append为切片添加数据时，当添加后的切片长度超过切片本来的容量时，切片容量会翻倍扩容
	// 即初始容量为10，添加数据后的切片长度为15，则容量为20，再次添加数据后切片长度为25，容量翻倍为40
	// 获取切片容量的函数：cap
	slice := make([]int, 10, 20)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) // 20

	slice1 := make([]int, 10)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // 10

	slice1 = append(slice1, 1, 2, 3, 4, 5)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // 20

	slice1 = append(slice1, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // 40

	slice1 = append(slice1, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // 80
}
