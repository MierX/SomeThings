/*
 * @Author MierX
 * @Title 23多个if和if_elseif的区别
 * @Date 2023-03-01 周一
 * @Time 15:34:42
 */

package main

import "fmt"

func main() {
	a := 10
	// 这种好
	if a != 10 {
		fmt.Println(1)
	} else if a == 10 {
		fmt.Println(a)
	} else {
		fmt.Println(3)
	}

	if a != 10 {
		fmt.Println(1)
	}

	if a == 10 {
		fmt.Println(a)
	}

	if a > 10 {
		fmt.Println(3)
	}
}
