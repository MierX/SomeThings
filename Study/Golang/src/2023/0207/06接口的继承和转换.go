/*
 * @Author MierX
 * @Title 06接口的继承和转换
 * @Date 2023-02-07 周一
 * @Time 17:45:12
 */

package main

import "fmt"

type Humaner interface { // 子集
	sayHi()
}

type Personnel interface { // 超集
	Humaner

	sing(string)
}

type Student2 struct {
	id  int
	age int
}

func (s *Student2) sayHi() {
	fmt.Printf("大家好，我的编号是%d，我的年龄是%d\n", s.id, s.age)
}

func (s *Student2) sing(name string) {
	fmt.Printf("我为大家唱一首：%s\n", name)
}

func main0601() {
	var h Humaner
	stu := Student2{12, 15}
	h = &stu
	h.sayHi()

	// 定义接口类型变量
	var p Personnel
	p = &stu
	p.sayHi()
	p.sing("传奇")
}

func main() {
	var h Humaner
	var p Personnel
	stu := Student2{12, 15}

	p = &stu
	h = p
	h.sayHi()

	h = &stu
	// 超集中，包含了所有子集的方法
	// 可以将超集赋值给子集 不能将子集赋值给超集
	// p = h // h 缺少了 p 里面的方法

}
