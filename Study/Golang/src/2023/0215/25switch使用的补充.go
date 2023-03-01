/*
 * @Author MierX
 * @Title 25switch使用的补充
 * @Date 2023-03-01 周一
 * @Time 15:44:04
 */

package main

import "fmt"

func main() {
	//支持一个初始化语句 初始化语句和变量本身以分号分割
	switch num := 1; num { //switch 后面写的是变量本身
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

	score := 85
	switch {
	case score > 90: //case后面可以放条件
		fmt.Println("A")
		break
	case score > 80:
		fmt.Println("B")
		break
	case score > 70:
		fmt.Println("C")
		break
	default:
		fmt.Println("D")
	}
}
