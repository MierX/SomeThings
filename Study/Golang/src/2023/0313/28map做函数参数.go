/*
 * @Author MierX
 * @Title 28map做函数参数
 * @Date 2023-03-21 周一
 * @Time 14:57:54
 */

package main

import "fmt"

func test(m map[int]string) {
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "mike", 2: "yo", 3: "go"}
	fmt.Println("m = ", m)

	test(m) //在函数内部删除某个key
	fmt.Println("m = ", m)
}
