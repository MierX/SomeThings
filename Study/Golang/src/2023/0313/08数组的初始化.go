/*
 * @Author MierX
 * @Title 08数组的初始化
 * @Date 2023-03-13 周一
 * @Time 15:51:08
 */

package main

import "fmt"

func main() {
	//声明定义同时赋值，叫初始化
	//1、全部初始化
	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	//2、部分初始哈，没有初始化的元素，自动赋值为0
	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)

	//3、指定某个元素初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Println("d = ", d)
}
