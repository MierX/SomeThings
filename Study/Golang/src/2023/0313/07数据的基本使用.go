/*
 * @Author MierX
 * @Title 07数据的基本使用
 * @Date 2023-03-13 周一
 * @Time 15:37:04
 */

package main

import "fmt"

func main() {
	//定义一个数组，[10]int和[5]int是不同类型
	//[数字]：作为数组元素个数
	var a [10]int
	var b [5]int

	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))

	//注意，定义数组时，指定的数组元素个数必须是常量

	//操作数组元素，从0开始，到len()-1，不对称元素，这个数字叫下标
	a[0] = 1
	i := 1
	a[i] = 2
}
