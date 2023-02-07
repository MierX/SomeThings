/*
 * @Author MierX
 * @Title 03匿名字段成员为指针
 * @Date 2023-02-07 周一
 * @Time 13:43:34
 */

package main

import "fmt"

type person1 struct {
	id   int
	name string
	age  int
	sex  string
}

type student1 struct {
	*person1

	score int
	class int
}

func main() {
	var stu student1
	stu.class = 301
	stu.score = 99

	// 需要对指针进行创建空间
	stu.person1 = new(person1)
	stu.name = "盖伦"
	stu.age = 30
	stu.sex = "男"
	stu.id = 1

	fmt.Println(stu)
	fmt.Printf("%s\n", stu.name)

	(*stu.person1).name = "诺手"
	fmt.Printf("%s\n", stu.name)
}
