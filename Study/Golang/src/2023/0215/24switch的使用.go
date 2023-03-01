/*
 * @Author MierX
 * @Title 24switch的使用
 * @Date 2023-03-01 周一
 * @Time 15:37:15
 */

package main

import "fmt"

func main() {
	var num int
	fmt.Printf("请按下楼层：")
	fmt.Scan(&num)

	switch num { //switch 后面写的是变量本身
	case 1:
		fmt.Println(1)
		break
	case 2:
		fmt.Println(2)
		break
	case 3:
		fmt.Println(3)
		break
	default:
		fmt.Println(num)
	}
}
