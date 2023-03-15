/*
 * @Author MierX
 * @Title 18切片和底层数组关系
 * @Date 2023-03-15 周一
 * @Time 16:30:56
 */

package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//新切片
	s1 := a[2:5]
	s1[1] = 666
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)

	//新切片
	s2 := a[2:7]
	s2[1] = 333
	fmt.Println("s2 = ", s2)
	fmt.Println("a = ", a)
}
