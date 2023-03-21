/*
 * @Author MierX
 * @Title 05结构体指针类型匿名字段
 * @Date 2023-03-21 周一
 * @Time 17:32:32
 */

package main

import "fmt"

type Person04 struct {
	name string
	sex  byte
	age  int
}

type Student04 struct {
	*Person04 //指针类型
	id        int
	addr      string
}

func main() {
	s1 := Student04{&Person04{"mike", 'm', 18}, 1, "bj"}
	fmt.Println(s1.name, s1.age, s1.sex, s1.addr, s1.id)

	//先定义变量
	var s2 Student04
	s2.Person04 = new(Person04) //分配空间
	s2.name = "jack"
	fmt.Println("s2 = ", s2)
}
