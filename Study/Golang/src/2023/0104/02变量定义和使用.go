/**
 * @Author MierX
 * @Title 02
 * @Date 2022-12-31 周六
 * @Time 02:26:15
 **/

package main

import (
	"fmt"
	"math"
)

func main() {
	// 变量的声明
	var num int

	fmt.Println(num)

	main001()

	main002()

	main003()
}

func main001() {

	/*
	 * 变量的定义
	 * var 声明变量
	 * num 变量名称
	 * int 变量类型：整型
	 * = 50 赋值变量
	 */
	var num int = 50

	// 表达式
	num += 25

	fmt.Println(num)
}

func main002() {
	var num int = 2

	var sum int = num * num * num * num * num * num * num * num * num * num

	fmt.Println(sum)
}

func main003() {
	var num float64 = 2

	var sum = math.Pow(num, 10)

	fmt.Println(sum)
}
