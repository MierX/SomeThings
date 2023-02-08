/*
 * @Author MierX
 * @Title 02panic异常
 * @Date 2023-02-08 周一
 * @Time 14:58:38
 */

package main

import (
	"fmt"
)

func test(i int) {
	var arr [10]int
	arr[i] = 123
	fmt.Println(arr)
}

func main0201() {
	fmt.Println("hello world1")
	fmt.Println("hello world2")
	// panic("hello world3")
	// println("hello world4")
	// println("hello world5")
}

func main() {
	test(10)
}
