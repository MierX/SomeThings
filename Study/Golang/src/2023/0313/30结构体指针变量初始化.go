/*
 * @Author MierX
 * @Title 30结构体指针变量初始化
 * @Date 2023-03-21 周一
 * @Time 15:11:20
 */

package main

import "fmt"

// Student01 定义一个结构体类型
type Student01 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//顺序初始化，每个成员必须初始化
	var p1 = &Student01{1, "mike", 'm', 18, "bj"}
	fmt.Println("p1 = ", p1)
	fmt.Println("*p1 = ", *p1)

	p2 := &Student01{name: "mike"}
	fmt.Println("p2 = ", p2)
	fmt.Println("*p2 = ", *p2)
}
