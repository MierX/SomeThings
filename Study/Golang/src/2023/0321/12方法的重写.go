/*
 * @Author MierX
 * @Title 12方法的重写
 * @Date 2023-03-24 周一
 * @Time 16:22:55
 */

package main

import "fmt"

type Person10 struct {
	name string
	sex  byte
	age  int
}

// PrintInfo Person10类型实现一个方法
func (tmp *Person10) PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d\n", tmp.name, tmp.sex, tmp.age)
}

// Student06 有个血酸，继承Person10字段，成员和方法都继承了
type Student06 struct {
	Person10 //匿名字段
	id       int
	addr     string
}

// PrintInfo Student06类型实现一个方法，与Person10的方法同名，这种方法叫重写
func (tmp *Student06) PrintInfo() {
	fmt.Println("Student06: tmp = ", tmp)
}

func main() {
	s := Student06{Person10{"mike", 'm', 18}, 666, "bj"}
	//就近原则：先找笨作用域的方法，找不到再用继承的
	s.PrintInfo()
	s.Person10.PrintInfo()
}
