/*
 * @Author MierX
 * @Title 03函数的嵌套调用.go
 * @Date 2023-01-18 周一
 * @Time 17:42:05
 */

package main

import "fmt"

// 函数参数传递时，如果有多个参数，中间用逗号分隔
func test(a int, b int) {
	test1(a, b)
}

func test1(a, b int) {
	fmt.Println(a + b)
}

func main0401() {
	a := 10
	b := 20
	test(a, b)
}

// 不定参函数的嵌套调用
func test2(arr ...int) {
	// 集合切片 前面是从哪个下标开始 后面是长度
	test3(arr[1:1]...)
}

func test3(arr ...int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	test2(1, 2, 3, 4)
}
