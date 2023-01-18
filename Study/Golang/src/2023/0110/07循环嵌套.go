/*
 * @Author MierX
 * @Title 07循环嵌套
 * @Date 2023-01-10 周二
 * @Time 23:05:31
 */

package main

import (
	"fmt"
	"time"
)

func main0701() {
	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			count++
			fmt.Println(i, j)
		}
	}
	fmt.Println(count)
}

func main0702() {
	for i := 0; i < 24; i++ {
		for j := 0; j < 60; j++ {
			for k := 0; k < 60; k++ {
				time.Sleep(time.Millisecond * 950)
				fmt.Println(i, j, k)
			}
		}
	}
}

func main0703() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
}

func main() {
	// 九九乘法口诀表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				// 跳出本次循环
				continue
			}
			fmt.Printf(" %d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
