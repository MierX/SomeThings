/*
 * @Author MierX
 * @Title 20append扩容特点
 * @Date 2023-03-15 周一
 * @Time 16:42:26
 */

package main

import "fmt"

func main() {
	//如果超过原来的容器，通常以2倍容量扩容
	s := make([]int, 0, 1) //容量为1
	oldCap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap: %d ===> %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}
