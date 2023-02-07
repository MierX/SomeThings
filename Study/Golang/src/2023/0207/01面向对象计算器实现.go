/*
 * @Author MierX
 * @Title 01面向对象计算器实现
 * @Date 2023-02-07 周一
 * @Time 16:44:11
 */

package main

import "fmt"

type Operate struct {
	num1 int
	num2 int
}

type Add struct {
	Operate
}

type Sub struct {
	Operate
}

func (a *Add) Result() int {
	return a.num1 + a.num2
}

func (s *Sub) Result() int {
	return s.num1 - s.num2
}

func main() {
	// 创建加法对象
	a := Add{Operate{10, 20}}
	rs := a.Result()
	fmt.Println(rs)

	s := Sub{Operate{10, 20}}
	rs1 := s.Result()
	fmt.Println(rs1)
}
