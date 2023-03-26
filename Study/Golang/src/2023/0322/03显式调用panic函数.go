/*
 * @Author MierX
 * @Title 03显式调用panic函数
 * @Date 2023-03-26 周一
 * @Time 11:59:07
 */

package main

import "fmt"

func testA() {
	fmt.Println("A")
}

func testB() {
	fmt.Println("B")
	//显式调用panic函数，导致程序中断
	panic("this is a panic testB")
}

func testC() {
	fmt.Println("C")
}

func main() {
	testA()
	testB()
	testC()
}
