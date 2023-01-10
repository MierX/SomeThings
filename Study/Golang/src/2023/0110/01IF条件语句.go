/*
 * @Author MierX
 * @Title 01IF条件语句
 * @Date 2023-01-10 周二
 * @Time 21:57:19
 */

package main

import "fmt"

func main() {
	// if 表达式 {
	//    代码体
	// } else if 表达式 {
	//    代码体
	// } else {
	//    代码体
	// }
	var score int
	fmt.Scan(&score)
	if score > 700 {
		fmt.Println("成功")
	} else if score > 600 {
		fmt.Println("快成功")
	} else {
		fmt.Println("失败")
	}
}
