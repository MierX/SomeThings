/*
 * @Author MierX
 * @Title 02接口定义和使用
 * @Date 2023-02-07 周一
 * @Time 16:55:51
 */

package main

import "fmt"

// 先定义接口 再根据接口实现功能

// Human 接口定义
// 接口定义的规则 接口定义的方法必须有实现
type Human interface {
	// SayHello 定义方法
	SayHello()
	Result(int, int) int
}

type Student struct {
	id  int
	age int
}

type Teacher struct {
	id       int
	age      int
	workTime int
}

func (s *Student) SayHello() {
	fmt.Printf("大家好，我的编号是%d，年龄是%d\n", s.id, s.age)
}

func (s *Student) Result(a, b int) int {
	return a - b
}

func (t Teacher) SayHello() {
	fmt.Printf("大家好，我的编号是%d，年龄是%d，我工作了%d年\n", t.id, t.age, t.workTime)
}

func main() {
	stu := Student{101, 18}
	stu.SayHello()

	tea := Teacher{1001, 35, 8}
	tea.SayHello()

	// 接口是一种数据类型 可以接收满足对象的信息
	// 接口是虚的 方法是实的
	// 接口定义规则 方法实现规则
	var i Human
	fmt.Printf("%T\n", i)
	// 将对象赋值给接口类型变量
	// 必须满足接口中方法的声明格式
	i = &stu
	i.SayHello()
}
