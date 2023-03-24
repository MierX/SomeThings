/*
 * @Author MierX
 * @Title 13方法值
 * @Date 2023-03-24 周一
 * @Time 16:27:39
 */

package main

import "fmt"

type Person11 struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值

// SetInfoValue 接收者为普通变量，非指针，值语义，一份拷贝
func (tmp Person11) SetInfoValue() {
	fmt.Printf("SetInfoValue: &tmp = %p, tmp = %v\n", &tmp, tmp)
}

// SetInfoPointer 接收者为指针变量，引用传递
func (tmp *Person11) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: &tmp = %p, tmp = %v\n", tmp, tmp)
}

func main() {
	p := Person11{"mike", 'm', 19}
	fmt.Printf("main: &p = %p, p = %v\n", &p, p)

	p.SetInfoPointer() //传统调用方式

	//保存方法入口地址
	pFunc := p.SetInfoPointer //这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者
	pFunc()

	vFunc := p.SetInfoValue
	vFunc()
}
