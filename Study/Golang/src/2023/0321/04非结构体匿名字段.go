/*
 * @Author MierX
 * @Title 04非结构体匿名字段
 * @Date 2023-03-21 周一
 * @Time 17:29:06
 */

package main

import "fmt"

type myStr string

type Person03 struct {
	name string
	sex  byte
	age  int
}

type Student03 struct {
	Person03 //只有类型，没有名字，匿名字段，继承了Person的成员
	int      //基础类型的同名字段
	myStr
}

func main() {
	s := Student03{Person03{"mike", 'm', 18}, 1, "bj"}
	fmt.Printf("s = %+v\n", s)
}
