/*
 * @Author MierX
 * @Title 03SWITCH控制语句
 * @Date 2023-01-10 周二
 * @Time 22:14:55
 */

package main

import "fmt"

func main0301() {
	// switch 变量 {
	// case 值1:
	// 	代码体
	// case 值2:
	// 	代码体
	// case 值3:
	// 	代码体
	// default:
	// 	代码体
	// }

	var score int
	fmt.Scan(&score)
	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9, 8:
		fmt.Println("B")
	case 7, 6:
		fmt.Println("C")
	case 5:
		fmt.Println("D")
	case 4, 3, 2:
		fmt.Println("E")
	default:
		fmt.Println("F")
	}
}

func main0302() {
	// switch {
	// case 表达式1:
	// 	代码体
	// case 表达式2:
	// 	代码体
	// case 表达式3:
	// 	代码体
	// default:
	// 	代码体
	// }
	var score int
	fmt.Scan(&score)
	switch {
	case score > 90:
		fmt.Println("A")
	case score <= 90 && score > 80:
		fmt.Println("B")
	case score <= 80 && score > 60:
		fmt.Println("C")
	case score <= 60:
		fmt.Println("D")
	}
}

func main0303() {
	var score int
	fmt.Scan(&score)
	switch score >= 60 {
	case true:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

func main0304() {
	// switch {
	// case 表达式1, 表达式2, 表达式3:
	// 	代码体
	// case 表达式4, 表达式5:
	// 	代码体
	// default:
	// 	代码体
	// }
	var y, m int
	fmt.Scan(&y, &m)
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 2:
		if (y%4 == 0 && y%100 != 0) || (y%400 == 0) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	case 4, 6, 9, 11:
		fmt.Println(30)
	default:
		fmt.Println("月份输入错误")
	}
}

func main() {
	// switch {
	// case 表达式1:fallthrough
	// case 表达式2:fallthrough
	// case 表达式3:
	// 	代码体
	// case 表达式4:fallthrough
	// case 表达式5:fallthrough
	// 	代码体
	// default:
	// 	代码体
	// }
	var y, m int
	fmt.Scan(&y, &m)
	switch m {
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 5:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 10:
		fallthrough
	case 12:
		fmt.Println(31)
	case 2:
		if (y%4 == 0 && y%100 != 0) || (y%400 == 0) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	case 4:
		fallthrough
	case 6:
		fallthrough
	case 9:
		fallthrough
	case 11:
		fmt.Println(30)
	default:
		fmt.Println("月份输入错误")
	}

}
