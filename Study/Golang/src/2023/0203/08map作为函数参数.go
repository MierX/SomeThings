/*
 * @Author MierX
 * @Title 08map作为函数参数
 * @Date 2023-02-05 周日
 * @Time 14:59:19
 */

package main

import "fmt"

func main() {
	m := map[int]string{1: "张三", 2: "李四", 3: "王五"}
	fmt.Println(m)
	// 引用传递
	demo(m)
	fmt.Println(m)
}

func demo(m map[int]string) {
	fmt.Println(m)
	m[4] = "赵六"
}
