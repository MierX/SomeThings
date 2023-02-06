/*
 * @Author MierX
 * @Title 07指针和切片
 * @Date 2023-02-06 周一
 * @Time 16:51:54
 */

package main

import "fmt"

func main0701() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 指针和切片建立关系
	p := &slice
	fmt.Printf("%p\n", p)
	fmt.Printf("%T\n", p)

	(*p)[0] = 222
	fmt.Println(slice)

	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	p := &s
	test3(p)
	fmt.Println(p)
}

func test3(s *[]int) {
	*s = append(*s, 6)
}
