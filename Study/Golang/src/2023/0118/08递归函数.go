/*
 * @Author MierX
 * @Title 08递归函数.go
 * @Date 2023-01-31 周一
 * @Time 16:42:45
 */

package main

import (
	"fmt"
	"strconv"
)

// 递归函数
// 没有出口的递归函数是死递归 可能导致程序崩溃
func demo4(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	n--
	demo4(n)
}

func main0801() {
	demo4(10)
}

// 计算一个数的阶乘

var sum1 = 1
var sum1str = ""

func demo5(n int) {
	if n == 1 {
		return
	}

	demo5(n - 1)
	if sum1str == "" {
		sum1str = strconv.Itoa(sum1) + " * " + strconv.Itoa(n)
	} else {
		sum1str = sum1str + " * " + strconv.Itoa(n)
	}
	sum1 *= n
}

func main() {
	demo5(6)
	fmt.Println(sum1str + " = " + strconv.Itoa(sum1))
}
