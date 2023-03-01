/*
 * @Author MierX
 * @Title 03不定参数类型
 * @Date 2023-03-01 周一
 * @Time 16:22:35
 */

package main

import "fmt"

func MyFunc05(a, b int) { //固定参数
	fmt.Println(a, b)
}

// MyFunc06 ...int类型这样的类型是不定参数，每个参数都是int类型
func MyFunc06(args ...int) { //传递的实参可以是0或多个
	fmt.Println(len(args))
	for i, arg := range args {
		fmt.Println(i, arg)
	}
}

func main() {
	MyFunc06(1, 2, 3)
}
