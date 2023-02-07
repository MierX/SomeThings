/*
 * @Author MierX
 * @Title 07方法的内存模型
 * @Date 2023-02-07 周一
 * @Time 15:16:59
 */

package main

import "fmt"

type student6 struct {
	name string
}

func main() {
	stu := student6{"aaa"}
	fmt.Println(stu)

	// 值传递
	stu.Print()
	fmt.Println(stu)

	// 引用传递
	stu.Print1()
	fmt.Println(stu)

}

func (s student6) Print() {
	s.name = "bbb"
}

// Print1 在方法的调用中 方法的接收者为指针类型
// 指针类型 和 普通类型 表示时相同对象的类型
func (s *student6) Print1() {
	s.name = "bbb"
}
