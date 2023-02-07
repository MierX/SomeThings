/*
 * @Author MierX
 * @Title 10方法重写
 * @Date 2023-02-07 周一
 * @Time 16:06:45
 */

package main

import "fmt"

type person8 struct {
	id  int
	age int
}

type student8 struct {
	person8
	class int
	score int
}

func (p *person8) SayHi() {
	fmt.Printf("编号：%d，年龄：%d\n", p.id, p.age)
}

// SayHi 子类对象和父类对象方法重名
func (s *student8) SayHi() {
	fmt.Printf("编号：%d，年龄：%d，班级：%d，成绩：%d\n", s.id, s.age, s.class, s.score)
}

func main() {
	stu := student8{person8{101, 18}, 202, 99}

	// 子类和父类有重名的方法时，采用就近原则，直接调用子类或最近一个继承类的方法
	// 如果子类与父类存在相同的方法，即子类对父类方法的重写
	stu.SayHi()

	// 调用父类的重名方法
	stu.person8.SayHi()

	fmt.Println(stu.SayHi)         // 0xcc0040
	fmt.Println(stu.person8.SayHi) // 0xcc00a0
}
