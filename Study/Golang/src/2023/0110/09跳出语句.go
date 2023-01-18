/*
 * @Author MierX
 * @Title 09跳出语句.go
 * @Date 2023-01-18 周一
 * @Time 14:52:04
 */

package main

import "fmt"

func main0901() {
	// break 结束整个循环体
	var i = 0
	for ; ; i++ {
		if i >= 5 {
			// 结束循环
			break
		}

		fmt.Println(i)
	}
}

func main0902() {
	// continue 结束当前一次循环，直接进入下次循环
	var i = 0
	for ; i <= 5; i++ {
		if i == 3 {
			// 结束循环
			continue
		}

		fmt.Println(i)
	}

}

func main() {
	// goto 跳至定位执行语句
	fmt.Println("1")
	fmt.Println("2")
	goto FLAG
	fmt.Println("3")
	fmt.Println("4")
FLAG:
	fmt.Println("5")
}
