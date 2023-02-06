/*
 * @Author MierX
 * @Title 02结构体作为函数参数
 * @Date 2023-02-05 周日
 * @Time 19:37:45
 */

package main

import "fmt"

type Student struct {
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main() {
	stu := Student{101, "喜羊羊", 6, "男", 100, "羊村"}
	// 结构体作为函数参数是值传递
	test(stu)
	fmt.Println(stu)

	m := map[int]Student{102: {101, "喜羊羊", 6, "男", 100, "羊村"}}
	// 结构体作为切片的元素，被传入函数是引用传递
	test1(m)
	fmt.Println(m)
}

func test(stu Student) {
	stu.name = "野猪佩奇"
	fmt.Println(stu)
}

func test1(m map[int]Student) {
	// m[102].name = "威震天" // cannot assign to struct field m[102].name in map
	stu := m[102]
	stu.name = "威震天"
	m[102] = stu
}
