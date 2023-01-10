/*
 * @Author MierX
 * @Title 02IF语句嵌套
 * @Date 2023-01-10 周二
 * @Time 22:06:11
 */

package main

import "fmt"

func main011001() {
	var score int
	fmt.Scan(&score)
	if score > 700 {
		fmt.Println("我要上清华")
		if score > 720 {
			fmt.Println("我可以选择热门专业")
		} else if score > 735 {
			fmt.Println("我可以选择精英专业")
		} else {
			fmt.Println("我只能选择冷门专业")
		}
	} else {
		fmt.Println("我上不了清华")
	}
}

func main() {
	// 案例：三只小猪称体重
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > b {
		if a > c {
			fmt.Println("a最重")
		} else {
			fmt.Println("c最重")
		}
	} else {
		if b > c {
			fmt.Println("b最重")
		} else {
			fmt.Println("c最重")
		}
	}
}
