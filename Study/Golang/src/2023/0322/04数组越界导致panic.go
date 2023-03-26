/*
 * @Author MierX
 * @Title 04数组越界导致panic
 * @Date 2023-03-26 周一
 * @Time 12:01:59
 */

package main

import "fmt"

func testA01() {
	fmt.Println("A")
}

func testB01(x int) {
	var a [10]int
	a[x] = 111 //当x为20的时候，导致数组越界，产生一个panic，导致程序崩溃
}

func testC01() {
	fmt.Println("C")
}

func main() {
	testA01()
	testB01(20)
	testC01()
}
