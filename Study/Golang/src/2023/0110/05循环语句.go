/*
 * @Author MierX
 * @Title 05循环语句
 * @Date 2023-01-10 周二
 * @Time 22:49:24
 */

package main

import "fmt"

func main0501() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
