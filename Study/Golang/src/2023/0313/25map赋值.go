/*
 * @Author MierX
 * @Title 25map赋值
 * @Date 2023-03-21 周一
 * @Time 14:48:18
 */

package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "yo"}
	//赋值，如果已经存在的key值，修改内容
	fmt.Println("m1 = ", m1)
	m1[1] = "c++"
	fmt.Println("m1 = ", m1)
	m1[3] = "go" //追加，map底层自动扩容，和append类似
	fmt.Println("m1 = ", m1)
}
