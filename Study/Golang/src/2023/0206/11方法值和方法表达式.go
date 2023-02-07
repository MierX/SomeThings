/*
 * @Author MierX
 * @Title 11方法值和方法表达式
 * @Date 2023-02-07 周一
 * @Time 16:14:21
 */

package main

import (
	"fmt"
)

type student9 struct {
	id  int
	age int
}

func (s *student9) PrintInfo() {
	fmt.Println(*s)
}

func test() {
	fmt.Println("hello world")
}

func test1() {
	fmt.Println("a")
}

func (s *student9) EditInfo(id, age int) {
	s.id = id
	s.age = age
}

func main1101() {
	stu := student9{101, 18}
	stu.PrintInfo()

	f := test
	fmt.Printf("%T\n", f) // func()
	f()

	s := stu.PrintInfo
	fmt.Printf("%T\n", s) // func()
	s()

	s = test1
	fmt.Printf("%T\n", s) // func()
	s()
}

func main() {
	stu := student9{101, 18}
	stu.EditInfo(202, 32)
	fmt.Println(stu)
}
