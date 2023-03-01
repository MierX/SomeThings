/*
 * @Author MierX
 * @Title 06多个返回值
 * @Date 2023-03-01 周一
 * @Time 16:48:17
 */

package main

import "fmt"

func MyFunc11() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}

func main() {
	a, b, c := MyFunc11()
	fmt.Println(a, b, c)
}
