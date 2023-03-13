/*
 * @Author MierX
 * @Title 01指针基本操作
 * @Date 2023-03-13 周一
 * @Time 15:08:04
 */

package main

import "fmt"

func main() {
	a := 10
	//每个变量有2层含义：变量的内存，变量的地址
	fmt.Printf("%d\n", a)  //变量的内存
	fmt.Printf("%v\n", &a) //变量的地址

	//保存某个变量的地址，需要指针类型 *int保存int的地址，**int保存*int地址
	//声明（定义），定义只是特殊的声明
	//定义一个指针变量p 类型为*int
	var p *int
	p = &a   //指针变量指向谁，就把谁的地址赋值给指针变量
	*p = 666 //*p操作的不是p的内存，是p所指向的内存（就是a）
	fmt.Printf("p = %v, a= %v\n", *p, a)
}
