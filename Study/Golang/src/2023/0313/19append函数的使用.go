/*
 * @Author MierX
 * @Title 19append函数的使用
 * @Date 2023-03-15 周一
 * @Time 16:38:48
 */

package main

import "fmt"

func main() {
	s1 := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	//在原切片的末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)
	s2 = append(s2, 4)
	s2 = append(s2, 5)
	s2 = append(s2, 6)
	fmt.Println("s2 = ", s2)
}
