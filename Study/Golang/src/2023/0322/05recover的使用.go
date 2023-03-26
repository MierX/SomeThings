/*
 * @Author MierX
 * @Title 05recover的使用
 * @Date 2023-03-26 周一
 * @Time 12:05:16
 */

package main

import "fmt"

func testA02() {
	fmt.Println("A")
}

func testB02(x int) {
	//设置recover
	defer func() {
		//recover()
		//fmt.Println(recover()) //可以打印panic的错误信息
		if err := recover(); err != nil { //产生了panic异常
			fmt.Println(err)
		}
	}() //调用此匿名函数
	var a [10]int
	a[x] = 111 //当x为20的时候，导致数组越界，产生一个panic，导致程序崩溃
	fmt.Println(a)
}

func testC02() {
	fmt.Println("C")
}

func main() {
	testA02()
	testB02(20)
	testC02()
}
