/*
 * @Author MierX
 * @Title 21if的使用
 * @Date 2023-03-01 周一
 * @Time 15:22:21
 */

package main

import "fmt"

func main() {
	s := "a"

	//if 条件语句 {代码体} 条件通常都是关系运算符
	if s == "a" {
		fmt.Println(s)
	}

	//if支持1个初始化语句
	if a := 10; a == 10 { //条件为真就执行
		fmt.Println(a)
	}
}
