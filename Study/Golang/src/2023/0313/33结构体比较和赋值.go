/*
 * @Author MierX
 * @Title 33结构体比较和赋值
 * @Date 2023-03-21 周一
 * @Time 15:21:27
 */

package main

import "fmt"

// Student04 定义一个结构体类型
type Student04 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	s1 := Student04{1, "mike", 'm', 10, "bj"}
	s2 := Student04{1, "mike", 'm', 10, "bj"}
	s3 := Student04{2, "mike", 'm', 10, "bj"}
	fmt.Println("s1 == s2", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)

	//同类型的2个结构体可以相互赋值
	var tmp Student04
	tmp = s3
	fmt.Println("tmp = ", tmp)
}
