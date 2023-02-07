/*
 * @Author MierX
 * @Title 06对象方法的创建和使用
 * @Date 2023-02-07 周一
 * @Time 14:41:47
 */

package main

import "fmt"

// type关键字：
// 1、定义函数类型 定义结构体名称
// 2、为已存在的数据类型起别名

type Int int

func main0601() {
	// 编译过程
	// 1、预处理 包展开 替换数据类型
	// 2、编译 如果代码有错会提示 如果没错会编译成汇编文件
	// 3、汇编 将汇编文件转成二进制文件
	// 4、链接 将支持的库连接到程序中 变成可执行程序

	// 类型别名在编译时进行转换
	var a Int
	a = 10
	fmt.Println(a)
	fmt.Println(a + 20)

	var b int
	b = 20
	// fmt.Println(a + b) 类型不匹配
	fmt.Println(a + Int(b))
}

func (a Int) add(b Int) Int {
	return a + b
}

func main0602() {
	var a Int = 10
	var b Int = 20

	// 想要使用方法 必须是相同类型的对象才可以
	sum := a.add(b)
	fmt.Println(sum)
}

type student5 struct {
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main() {
	// 创建对象
	stu := student5{"贾政", 58, "男", 80, "贾府"}
	stu.Print()

	stu1 := student5{"贾政", 58, "男", 99, "贾府"}
	stu1.Print()

	Print()

	// 打印函数在代码区的地址
	fmt.Println(stu.Print)
	fmt.Println(Print)
}

// 结构体类型可以作为对象类型
// 创建对象方法

func (s student5) Print() {
	fmt.Printf("姓名：%s\n", s.name)
	fmt.Printf("年龄：%d\n", s.age)
	fmt.Printf("性别：%s\n", s.sex)
	fmt.Printf("成绩：%d\n", s.score)
	fmt.Printf("地址：%s\n", s.addr)
}

func Print() {
	fmt.Println("hello world")
}
