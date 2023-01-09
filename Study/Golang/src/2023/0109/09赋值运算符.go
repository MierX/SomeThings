/*
 * @Author MierX
 * @Title 09赋值运算符.go
 * @Date 2023-01-09 周一
 * @Time 14:48:57
 */

package main

import "fmt"

func main() {
	a := 10
	b := 20

	c := a + b
	fmt.Println(c)
	c += 20
	fmt.Println(c)
	c -= 20
	fmt.Println(c)
	c *= 20
	fmt.Println(c)
	c /= 20
	fmt.Println(c)
	c %= 8
	fmt.Println(c)
}
