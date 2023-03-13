/*
 * @Author MierX
 * @Title 15切片的长度和容量
 * @Date 2023-03-13 周一
 * @Time 17:05:29
 */

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 0, 0}
	s := a[0:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s)) //容量
}
