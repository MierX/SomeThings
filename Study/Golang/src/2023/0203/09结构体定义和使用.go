/*
 * @Author MierX
 * @Title 09结构体定义和使用
 * @Date 2023-02-05 周日
 * @Time 15:07:35
 */

package main

import "fmt"

// 结构体敌营在函数外部进行
// 定义： type 结构体名 struct {结构体成员列表}
// 结构体是一种数据类型
type student struct {
	// 成员名 数据类型
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main() {
	var stu student
	fmt.Println(stu)
	fmt.Printf("%T\n", stu)

	// 为结构体成员赋值
	stu.id = 1
	stu.name = "张三"
	stu.age = 18
	stu.sex = "男"
	stu.score = 100
	stu.addr = "中国"
	fmt.Println(stu)
	fmt.Println(stu.name)

	var stu1 = student{1, "李四", 17, "男", 95, "中国"}
	fmt.Println(stu1)

	m := make(map[int]student)
	m[1] = stu
	m[2] = stu1
	fmt.Println(m)
	fmt.Println(m[1].name)

	arr := make([]student, 0)
	arr = append(arr, stu, stu1)
	fmt.Println(arr)
	fmt.Println(arr[0].name)
}
