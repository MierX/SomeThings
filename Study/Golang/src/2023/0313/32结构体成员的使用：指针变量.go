/*
 * @Author MierX
 * @Title 32结构体成员的使用：指针变量
 * @Date 2023-03-21 周一
 * @Time 15:16:28
 */

package main

import "fmt"

// Student03 定义一个结构体类型
type Student03 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//1、指针有合法指向后，才操作成员
	//先定义一个普通结构体变量
	var s Student03
	//在定义一个指针变量，保存s的地址
	var p1 *Student03
	p1 = &s

	//通过指针操作成员 p1.id 和 *(p1).id 完全等价，只能使用点运算符
	p1.id = 18
	(*p1).name = "mike"
	p1.sex = 'm'
	fmt.Println("p1 = ", p1)

	//2、通过new来申请一个结构体
	p2 := new(Student03)
	p2.id = 1
	p2.name = "mike"
	fmt.Println("p2 = ", p2)
}
