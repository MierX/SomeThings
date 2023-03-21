/*
 * @Author MierX
 * @Title 03同名字段
 * @Date 2023-03-21 周一
 * @Time 17:22:52
 */

package main

import "fmt"

type Person02 struct {
	name string
	sex  byte
	age  int
}

type Student02 struct {
	Person02 //只有类型，没有名字，匿名字段，继承了Person的成员
	id       int
	addr     string
	name     string
}

func main() {
	//声明
	var s Student02
	s.name = "mike"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	s.id = 1
	fmt.Printf("s = %+v\n", s)
}
