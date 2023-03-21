/*
 * @Author MierX
 * @Title 31结构体成员的使用：普通变量
 * @Date 2023-03-21 周一
 * @Time 15:14:44
 */

package main

import "fmt"

// Student02 定义一个结构体类型
type Student02 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//定义一个结构体普通变量
	var s Student02

	//操作成员，需要使用点运算符
	s.id = 1
	s.name = "mike"
	fmt.Println("s = ", s)
}
