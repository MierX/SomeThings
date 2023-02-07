/*
 * @Author MierX
 * @Title 04多重继承
 * @Date 2023-02-07 周一
 * @Time 14:16:21
 */

package main

import "fmt"

type person2 struct {
	name string
}

type person3 struct {
	id   int
	addr string
}

type student2 struct {
	person2
	person3

	class int
	score int
}

type student3 struct {
	student2

	sex string
}

func main0401() {
	var stu student2
	stu.id = 1001
	stu.name = "亚索"
	stu.score = -5
	stu.class = 300
	stu.addr = "瓦罗兰大陆"
	fmt.Println(stu)

	stu1 := student2{person2{"压缩"}, person3{1, "aa"}, 101, 99}
	fmt.Println(stu1)
}

func main() {
	var stu student3
	stu.id = 1001
	stu.name = "亚索"
	stu.score = -5
	stu.class = 300
	stu.addr = "瓦罗兰大陆"
	stu.sex = "男"
	fmt.Println(stu)
}
