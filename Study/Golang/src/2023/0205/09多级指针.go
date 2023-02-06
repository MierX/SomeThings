/*
 * @Author MierX
 * @Title 09多级指针
 * @Date 2023-02-06 周一
 * @Time 17:43:36
 */

package main

import "fmt"

func main0901() {
	a := 10

	// 一级指针 指向 变量的地址
	p := &a

	// 二级指针 指针 一级指针的地址
	p1 := &p

	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", p1)

	// 通过二级指针间接修改变量的值
	**p1 = 100
	fmt.Println(a)
}

func main() {
	a := 10
	p := &a
	var pp **int
	pp = &p

	var ppp = &pp

	pppp := &ppp
	fmt.Println(pppp)
}
