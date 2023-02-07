/*
 * @Author MierX
 * @Title 02匿名字段同名成员
 * @Date 2023-02-07 周一
 * @Time 13:30:26
 */

package main

import "fmt"

type person struct {
	id   int
	name string
	age  int
}

type student struct {
	person
	name  string
	class int
	score int
}

func main() {
	var stu student
	stu.score = 100
	stu.class = 301
	stu.id = 1002
	stu.age = 18
	// 有相同的成员名时，赋值并不会报错，但只会为直接的成员赋值
	stu.name = "瞎子"
	stu.person.name = "瞎子"
	fmt.Println(stu)
}
