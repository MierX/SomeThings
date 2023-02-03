/*
 * @Author MierX
 * @Title 09切片的截取.go
 * @Date 2023-02-03 周一
 * @Time 15:36:51
 */

package main

import "fmt"

func main0901() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 截取切片以下标为主，方法：切片变量[切片起始位置下标（含）:切片结束位置下标（不含）：容量母值（新切片的容量由 容量母值减去起始位置下标 决定， 且不能小于切片长度）]
	s := slice[3:7:7]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// 未定义容量母值，则容量由原切片容量减去切片起始下标决定
	s1 := slice[3:7]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	// 未定义结束位置则从起始位置截取切片全部
	s2 := slice[3:]
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	// 未定义起始位置则默认起始位置下标为0
	s3 := slice[:5:5]
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	// 未定义起始位置和结束位置则相当于复制
	s4 := slice[:]
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := slice[2:5]
	fmt.Println(s)     // [2 3 4]
	fmt.Println(slice) // [0 1 2 3 4 5 6 7 8 9]

	// 截取后的切片，仍然与原切片共用某块内存地址（引用），故新切片的数据发生改变后，原切片也会改变
	s[1] = 123
	fmt.Println(s)     // [2 123 4]
	fmt.Println(slice) // [0 1 2 123 4 5 6 7 8 9]

	s = append(s, 999)
	fmt.Println(s)     // [2 123 4 999]
	fmt.Println(slice) // [0 1 2 123 4 999 6 7 8 9]
}
