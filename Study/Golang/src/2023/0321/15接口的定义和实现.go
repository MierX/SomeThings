/*
 * @Author MierX
 * @Title 15接口的定义和实现
 * @Date 2023-03-24 周一
 * @Time 16:39:14
 */

package main

import "fmt"

// Humaner 定义接口类型
type Humaner interface {
	//方法，只有生命，没有实现，由别的类型（自定义类型）实现
	sayHi()
}

type Student07 struct {
	name string
	id   int
}

// Student07实现了sayHi方法
func (tmp *Student07) sayHi() {
	fmt.Printf("Student07[%s, %d] sayHi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

// Teacher实现了sayHi方法
func (tmp *Teacher) sayHi() {
	fmt.Printf("Teacher[%s, %s] sayHi\n", tmp.addr, tmp.group)
}

type MyStr string

// MyStr实现了sayHi方法
func (tmp *MyStr) sayHi() {
	fmt.Printf("MyStr[%s] sayHi\n", *tmp)
}

// WhoSayHi 定义一个普通函数，函数的参数为接口类型
func WhoSayHi(i Humaner) { //只有一个函数，可以有不同表现，就是多态
	i.sayHi()
}

func main() {
	s := &Student07{"mike", 666}
	t := &Teacher{"bj", "go"}
	var str MyStr = "hi"

	//调用同一函数，不同表现，多态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, humaner := range x {
		humaner.sayHi()
	}
}

func main01() {
	//定义接口类型的变量
	var i Humaner

	//只是实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值
	s := &Student07{"mike", 666}
	i = s
	i.sayHi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayHi()

	var str MyStr = "hi"
	i = &str
	i.sayHi()
}
