/*
 * @Author MierX
 * @Title 29结构体普通变量初始化
 * @Date 2023-03-21 周一
 * @Time 15:04:55
 */

package main

import "fmt"

// Student 定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//顺序初始化，每个成员必须初始化
	var s1 = Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("s1 = ", s1)

	//指定成员初始化，没有初始化的成员自动赋值为0
	s2 := Student{name: "mike"}
	fmt.Println("s2 = ", s2)
}
