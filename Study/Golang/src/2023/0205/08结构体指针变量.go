/*
 * @Author MierX
 * @Title 08结构体指针变量
 * @Date 2023-02-06 周一
 * @Time 17:22:28
 */

package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main0801() {
	per := Person{101, "张三丰", "男", 100, "武当山"}
	fmt.Printf("%p\n", &per)

	// 定义指针接收结构体变量地址
	p := &per
	fmt.Printf("%T\n", p)

	// 通过指针间接修改结构体成员的值 通过括号提高*运算符级别
	(*p).name = "张君宝"

	// 指针可以直接操作结构体成员
	p.age = 140
	fmt.Println(per)
}

// 将结构体指针作为函数参数
func test4(p *Person) {
	p.addr = "中国"
}

func main0802() {
	per := Person{101, "张三丰", "男", 100, "武当山"}
	p := &per
	test4(p)
	fmt.Println(per)
}

func main0803() {
	arr := [3]Person{
		{101, "张三丰", "男", 100, "武当山"},
		{101, "张三丰", "男", 100, "武当山"},
		{101, "张三丰", "男", 100, "武当山"},
	}
	p := &arr
	fmt.Printf("%p\n", p)
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}

	p[0].age = 40
	fmt.Println(arr)
}

func main() {
	m := make(map[string]*[3]Person)

	fmt.Printf("%T\n", m)
	m["a"] = new([3]Person)
	m["a"] = &[3]Person{
		{101, "张三丰", "男", 100, "武当山"},
		{101, "张三丰", "男", 100, "武当山"},
		{101, "张三丰", "男", 100, "武当山"},
	}
	fmt.Println(m)
	for s, i := range m {
		fmt.Println(s, i)
		fmt.Println(s, *i)
	}
}
