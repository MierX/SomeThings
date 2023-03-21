/*
 * @Author MierX
 * @Title 34结构体作为函数参数
 * @Date 2023-03-21 周一
 * @Time 15:24:46
 */

package main

import "fmt"

// Student05 定义一个结构体类型
type Student05 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test01(s Student05) {
	s.id = 666
	fmt.Println("test01: ", s)
}

func test02(s *Student05) {
	s.id = 666
	fmt.Println("test01: ", s)
}

func main() {
	s := Student05{1, "mike", 'm', 19, "bj"}
	test01(s) //值传递，形参无法改实参
	fmt.Println("main: ", s)
	test02(&s) //值传递，形参无法改实参
	fmt.Println("main: ", s)
}
