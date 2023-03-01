/*
 * @Author MierX
 * @Title 26for的使用
 * @Date 2023-03-01 周一
 * @Time 15:49:23
 */

package main

import "fmt"

func main() {
	//for 初始化条件；判断条件；条件变化{代码体}
	sum := 0
	//1、初始化条件：i := 1
	//2、判断条件：i <= 100
	//3、条件变化：i++
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
