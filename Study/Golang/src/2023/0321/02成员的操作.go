/*
 * @Author MierX
 * @Title 02成员的操作
 * @Date 2023-03-21 周一
 * @Time 17:19:44
 */

package main

import "fmt"

type Person01 struct {
	name string
	sex  byte
	age  int
}

type Student01 struct {
	Person01 //只有类型，没有名字，匿名字段，继承了Person的成员
	id       int
	addr     string
}

func main() {
	var s1 = Student01{Person01{"mike", 'm', 18}, 1, "bj"}
	s1.name = "jack"
	s1.Person01 = Person01{"mier", 'w', 19}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
}
