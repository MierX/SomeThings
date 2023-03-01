/*
 * @Author MierX
 * @Title 22if_elseif_else的使用
 * @Date 2023-03-01 周一
 * @Time 15:26:46
 */

package main

import "fmt"

func main() {
	a := 10
	if a != 10 {
		fmt.Println(1)
	} else if a == 10 {
		fmt.Println(a)
	} else {
		fmt.Println(3)
	}
}
