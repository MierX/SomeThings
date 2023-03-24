/*
 * @Author MierX
 * @Title 07为结构体类型添加方法
 * @Date 2023-03-24 周一
 * @Time 15:25:22
 */

package main

import "fmt"

type Person05 struct {
	name string
	sex  byte
	age  int
}

// PrintInfo 带有接收者的函数叫方法
func (tmp Person05) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

// SetInfo 通过一个函数，给成员赋值
func (tmp *Person05) SetInfo(n string, s byte, a int) {
	tmp.name = n
	tmp.sex = s
	tmp.age = a
}

// 只要是接收者类型不一样，方法可以同名，是不同的方法，不会出现重复定义函数的错误
type long01 int

func (tmp long01) test() {

}

type char byte

func (tmp char) test() {

}

func (tmp *long01) test02() {

}

type pointer *int

//pointer为接收者类型，它本身不能是指针类型
//func (tmp pointer) test() {
//
//}

func main() {
	//定义同时初始化
	p := Person05{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person05
	(&p2).SetInfo("yo", 'f', 32)
	p2.PrintInfo()
}
