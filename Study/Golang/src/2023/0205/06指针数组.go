/*
 * @Author MierX
 * @Title 06指针数组
 * @Date 2023-02-06 周一
 * @Time 16:34:51
 */

package main

import "fmt"

func main0601() {
	// 定义普通数组
	// var arr [10]int
	// 定义包含指针元素的数组 指针数组
	// var arr [3]*int
	a := 10
	b := 20
	c := 30
	arr := [3]*int{&a, &b, &c}
	fmt.Println(arr)
	*(arr[1]) = 200
	fmt.Println(b)
	fmt.Println(*arr[1])
	*arr[0] = 100
	fmt.Println(a)
}

func main() {
	a := [3]int{0, 1, 2}
	b := [3]int{3, 4, 5}
	c := [3]int{6, 7, 8}
	arr := [3]*[3]int{&a, &b, &c}
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	fmt.Printf("%p\n", &arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(*arr[i])
		for j := 0; j < len(*arr[i]); j++ {
			fmt.Printf("%d \n", (*arr[i])[j])
		}
	}

	(*arr[0])[0] = 555
	fmt.Println(a)
}
