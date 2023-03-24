/*
 * @Author MierX
 * @Title 10普通变量的方法集
 * @Date 2023-03-24 周一
 * @Time 16:16:28
 */

package main

import "fmt"

type Person08 struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值

// SetInfoValue 接收者为普通变量，非指针，值语义，一份拷贝
func (tmp Person08) SetInfoValue() {
	fmt.Printf("SetInfoValue &tmp = %p\n", &tmp)
}

// SetInfoPointer 接收者为指针变量，引用传递
func (tmp *Person08) SetInfoPointer() {
	fmt.Printf("SetInfoPointer tmp = %p\n", tmp)
}

func main() {
	//假如结构体变量是一个指针变量，他能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := Person08{"mike", 'm', 18}
	p.SetInfoValue()
	(&p).SetInfoPointer()

	//内部做的转换，先把（取地址p）&p转成p后调用
	(&p).SetInfoValue()

	//内部做的转换，先把p转成（取地址p）&p后调用
	p.SetInfoPointer()
}
