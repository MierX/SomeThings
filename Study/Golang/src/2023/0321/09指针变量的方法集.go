/*
 * @Author MierX
 * @Title 09指针变量的方法集
 * @Date 2023-03-24 周一
 * @Time 16:08:25
 */

package main

import "fmt"

type Person07 struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值

// SetInfoValue 接收者为普通变量，非指针，值语义，一份拷贝
func (tmp Person07) SetInfoValue() {
	fmt.Printf("SetInfoValue &tmp = %p\n", &tmp)
}

// SetInfoPointer 接收者为指针变量，引用传递
func (tmp *Person07) SetInfoPointer() {
	fmt.Printf("SetInfoPointer tmp = %p\n", tmp)
}

func main() {
	//假如结构体变量是一个指针变量，他能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person07{"mike", 'm', 18}
	p.SetInfoPointer()
	(*p).SetInfoValue()

	//内部做的转换，先把*p转成指针p后调用
	(*p).SetInfoPointer()

	//内部做的转换，先把指针p转成*p后调用
	p.SetInfoValue() //等同于 (*p).SetInfoValue()
}
