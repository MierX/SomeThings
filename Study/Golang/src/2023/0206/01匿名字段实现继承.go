/*
 * @Author MierX
 * @Title 01匿名字段实现继承
 * @Date 2023-02-06 周一
 * @Time 18:16:28
 */

package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
	sex  string
}

type Student struct {
	// 将另一个结构体作为结构体的成员 就是 继承
	// per Person // 非匿名字段
	// 匿名字段继承结构体
	Person
	id    int
	class int
	score int
}

func main0101() {
	// 创建对象
	var stu Student

	stu.Person.name = "张三"
	stu.Person.id = 1001
	stu.Person.age = 1001
	stu.sex = "女"
	stu.id = 2002
	stu.class = 302
	stu.score = 99

	fmt.Println(stu)
}

func main() {
	var stu = Student{Person{1001, "东方不败", 35, "不详"}, 2002, 302, 100}
	fmt.Printf("姓名：%s\n", stu.name)
	fmt.Printf("性别：%s\n", stu.sex)
	fmt.Printf("年龄：%d\n", stu.age)
	fmt.Printf("编号：%d\n", stu.id)
	fmt.Printf("班级：%d\n", stu.class)
	fmt.Printf("份数：%d\n", stu.score)
}
