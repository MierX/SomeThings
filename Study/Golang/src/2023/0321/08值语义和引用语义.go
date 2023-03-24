/*
 * @Author MierX
 * @Title 08值语义和引用语义
 * @Date 2023-03-24 周一
 * @Time 15:57:35
 */

package main

import "fmt"

type Person06 struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值

// SetInfoValue 接收者为普通变量，非指针，值语义，一份拷贝
func (tmp Person06) SetInfoValue(n string, s byte, a int) {
	tmp.name = n
	tmp.sex = s
	tmp.age = a
	fmt.Printf("SetInfoValue &tmp = %p\n", &tmp)
}

// SetInfoPointer 接收者为指针变量，引用传递
func (tmp *Person06) SetInfoPointer(n string, s byte, a int) {
	tmp.name = n
	tmp.sex = s
	tmp.age = a
	fmt.Printf("SetInfoPointer tmp = %p\n", tmp)
}

func main() {
	s1 := Person06{"go", 'm', 32}
	fmt.Printf("&s1 = %p\n", &s1)

	//值语义
	s1.SetInfoValue("mike", 'm', 18)
	fmt.Println("s1 = ", s1)

	//引用语义
	(&s1).SetInfoPointer("mike", 'm', 18)
	fmt.Println("s1 = ", s1)
}
