/*
 * @Author MierX
 * @Title 01匿名字段初始化
 * @Date 2023-03-21 周一
 * @Time 17:09:52
 */

package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
}

func main() {
	//顺序初始化
	var s1 = Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println("s1 = ", s1)

	//自动推导类型
	s2 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	//%+v,显示更详细
	fmt.Printf("s2 = %+v\n", s2)

	//指定成员初始化
	s3 := Student{id: 1}
	fmt.Println("s3 = ", s3)
	s4 := Student{Person: Person{name: "mike"}}
	fmt.Println("s4 = ", s4)
}
