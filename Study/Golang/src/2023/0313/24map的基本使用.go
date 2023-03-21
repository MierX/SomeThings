/*
 * @Author MierX
 * @Title 24map的基本使用
 * @Date 2023-03-21 周一
 * @Time 14:40:12
 */

package main

import "fmt"

func main() {
	//定义一个变量 类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len，没有cap
	fmt.Println("len(m1) = ", len(m1))

	//可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len(m2) = ", len(m2))

	//可以通过make创建，可以指定长度
	m3 := make(map[int]string, 2)
	m3[1] = "mike"
	m3[2] = "ok"
	m3[3] = "c++"
	fmt.Println("m3 = ", m3)
	fmt.Println("len(m2) = ", len(m3))

	//键值是唯一的
	m4 := map[int]string{1: "mike", 2: "ok", 3: "c++"}
	fmt.Println("m4 = ", m4)
}
