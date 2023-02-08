/*
 * @Author MierX
 * @Title 03延迟调用defer
 * @Date 2023-02-08 周一
 * @Time 15:56:44
 */

package main

import "fmt"

func main0301() {
	fmt.Println("hello world1")
	fmt.Println("hello world2")

	// defer是栈，先进后出 后进先出 同一个函数下，有多个延迟调用，先调用最后一个
	defer fmt.Println("hello world3")
	defer fmt.Println("hello world4")
	fmt.Println("hello world5")
}

func main() {
	a := 10
	b := 20
	f := func() {
		fmt.Println(a + b)
	}
	n := func(a, b int) {
		fmt.Println(a + b)
	}
	defer f()
	defer n(a, b)

	a = 123
	b = 321
}
