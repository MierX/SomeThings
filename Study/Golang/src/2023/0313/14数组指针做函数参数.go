/*
 * @Author MierX
 * @Title 14数组指针做函数参数
 * @Date 2023-03-13 周一
 * @Time 16:56:22
 */

package main

import "fmt"

// a指向实现数组a，它是指向数组，是数组指针
func modify1(a *[5]int) {
	(*a)[0] = 666
	fmt.Println("modify:a = ", *a)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify1(&a) //数组地址传递过去
	fmt.Println("main:a = ", a)
}
