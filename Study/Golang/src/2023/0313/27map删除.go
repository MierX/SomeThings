/*
 * @Author MierX
 * @Title 27map删除
 * @Date 2023-03-21 周一
 * @Time 14:56:09
 */

package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "yo", 3: "go"}
	fmt.Println("m = ", m)
	delete(m, 1) //删除key为1的内容
	fmt.Println("m = ", m)
}
