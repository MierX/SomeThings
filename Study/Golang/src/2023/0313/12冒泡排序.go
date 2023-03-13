/*
 * @Author MierX
 * @Title 12冒泡排序
 * @Date 2023-03-13 周一
 * @Time 16:20:40
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var a [10]int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println("a = ", a)

	//冒泡排序
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("a = ", a)
}
