/*
 * @Author MierX
 * @Title 14方法表达式
 * @Date 2023-03-24 周一
 * @Time 16:34:26
 */

package main

import "fmt"

type Person12 struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值

// SetInfoValue 接收者为普通变量，非指针，值语义，一份拷贝
func (tmp Person12) SetInfoValue() {
	fmt.Printf("SetInfoValue: &tmp = %p, tmp = %v\n", &tmp, tmp)
}

// SetInfoPointer 接收者为指针变量，引用传递
func (tmp *Person12) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: &tmp = %p, tmp = %v\n", tmp, tmp)
}

func main() {
	p := Person12{"mike", 'm', 19}
	fmt.Printf("main: &p = %p, p = %v\n", &p, p)

	//方法值： pFunc := p.SetInfoPointer 隐藏了接收者

	//方法表达式
	f := (*Person12).SetInfoPointer
	f(&p) //显式 把接收者传递过去

	f2 := Person12.SetInfoValue
	f2(p)
}
