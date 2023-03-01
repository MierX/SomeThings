/*
 * @Author MierX
 * @Title 28break和continue的区别
 * @Date 2023-03-01 周一
 * @Time 15:57:50
 */

package main

import "fmt"

func main() {
	i := 0
	for { //for 后面不写任何东西，这个循环条件永远为真 死循环
		if i == 0 || i == 10 {
			i++
			continue //跳出本次循环进入下一次循环
		} else if i > 20 {
			break //结束循环体
		} else {
			fmt.Println(i)
			i++
		}
	}
}
