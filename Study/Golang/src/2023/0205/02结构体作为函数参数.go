/*
 * @Author MierX
 * @Title 02结构体作为函数参数
 * @Date 2023-02-05 周日
 * @Time 19:37:45
 */

package main

import "fmt"

type Student struct {
	name  string
	sex   string
	age   int
	score int
	addr  string
}

func main() {
	stu := Student{"喜羊羊", "男", 8, 100, "羊村"}
	fmt.Println(stu)
	test(stu)
	fmt.Println(stu)
}

func test(stu Student) {
	stu.name = "沸羊羊"
}
