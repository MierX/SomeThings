/*
 * @Author MierX
 * @Title 04多态的实现
 * @Date 2023-02-07 周一
 * @Time 17:12:06
 */

package main

import "fmt"

// 先定义接口 再根据接口实现功能

// Human 接口定义
// 接口定义的规则 接口定义的方法必须有实现
type Human1 interface {
	// SayHello 定义方法
	SayHello()
	Result(int, int) int
}

type Student1 struct {
	id  int
	age int
}

type Teacher1 struct {
	id       int
	age      int
	workTime int
}

func (s *Student1) SayHello() {
	fmt.Printf("大家好，我的编号是%d，年龄是%d\n", s.id, s.age)
}

func (s *Student1) Result(a, b int) int {
	return a - b
}

func (t Teacher1) SayHello() {
	fmt.Printf("大家好，我的编号是%d，年龄是%d，我工作了%d年\n", t.id, t.age, t.workTime)
}

func main() {
	stu := Student1{1, 18}

	// 调用多态函数
	SayHi(&stu)
}

func SayHi(h Human1) {
	h.SayHello()
}
