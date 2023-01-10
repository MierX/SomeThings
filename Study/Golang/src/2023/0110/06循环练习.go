/*
 * @Author MierX
 * @Title 06循环练习
 * @Date 2023-01-10 周二
 * @Time 22:57:14
 */

package main

import "fmt"

func main0601() {
	// 抽7
	for i := 1; i <= 100; i++ {
		if i%7 == 0 || i%10 == 7 || i/10 == 7 {
			fmt.Println(i)
		}
	}
}

func main() {
	// 水仙花数
	for i := 100; i < 1000; i++ {
		a := i / 100
		b := i / 10 % 10
		c := i % 10
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}
