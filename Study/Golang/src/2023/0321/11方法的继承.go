/*
 * @Author MierX
 * @Title 11方法的继承
 * @Date 2023-03-24 周一
 * @Time 16:19:01
 */

package main

import "fmt"

type Person09 struct {
	name string
	sex  byte
	age  int
}

// PrintInfo Person09类型实现一个方法
func (tmp *Person09) PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d\n", tmp.name, tmp.sex, tmp.age)
}

// Student05 有个血酸，继承Person09字段，成员和方法都继承了
type Student05 struct {
	Person09 //匿名字段
	id       int
	addr     string
}

func main() {
	s := Student05{Person09{"mike", 'm', 18}, 666, "bj"}
	s.PrintInfo()
}
