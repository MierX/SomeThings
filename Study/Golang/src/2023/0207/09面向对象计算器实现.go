/*
 * @Author MierX
 * @Title 09面向对象计算器实现
 * @Date 2023-02-08 周一
 * @Time 14:08:52
 */

package main

import "fmt"

// Operate1 定义接口
type Operate1 interface {
	// Result 方法声明
	Result() int
}

// Opera 父类
type Opera struct {
	num1 int
	num2 int
}

// Add1 加法子类
type Add1 struct {
	Opera
}

// Result 加法子类的方法
func (a *Add1) Result() int {
	return a.num1 + a.num2
}

// Sub1 减法子类
type Sub1 struct {
	Opera
}

// Result 减法子类的方法
func (s *Sub1) Result() int {
	return s.num1 - s.num2
}

// Mlt 乘法子类
type Mlt struct {
	Opera
}

// Result 乘法子类的方法
func (m *Mlt) Result() int {
	return m.num1 * m.num2
}

// Result 定义多态函数
func Result(o Operate1) int {
	value := o.Result()
	// fmt.Println(value)
	return value
}

// Fac 工程类 工厂模式（一种设计模式） 本身也是一个结构体 空的
type Fac struct {
}

// Result 绑定工厂类的方法
func (f *Fac) Result(num1, num2 int, ch string) (result int) {
	switch ch {
	case "+":
		var a Add1
		a.num1, a.num2 = num1, num2
		result = Result(&a)
	case "-":
		var s Sub1
		s.num1, s.num2 = num1, num2
		result = Result(&s)
	case "*":
		var m Mlt
		m.num1, m.num2 = num1, num2
		result = Result(&m)
	}
	return
}

// 接口调用
func main0901() {
	var o Operate1
	var a = Add1{Opera{10, 20}}

	o = &a
	value := o.Result()
	fmt.Println(value)
}

// 多态调用
func main0902() {
	var a = Add1{Opera{10, 20}}
	Result(&a)

	var s = Sub1{Opera{10, 20}}
	Result(&s)
}

// 通过设计模式（工厂模式）调用
func main() {
	// 创建工厂对象
	var f Fac
	value := f.Result(10, 20, "*")
	fmt.Println(value)
}
