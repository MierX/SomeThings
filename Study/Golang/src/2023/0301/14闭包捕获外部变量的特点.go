/*
 * @Author MierX
 * @Title 14闭包捕获外部变量的特点
 * @Date 2023-03-09 周一
 * @Time 15:11:44
 */

package main

import "fmt"

func main() {
	a := 10
	str := "a"

	fmt.Printf("外部：a = %d, b = %s\n", a, str)

	func() {
		a = 666
		str = "go"
		fmt.Printf("内部：a = %d, b = %s\n", a, str)
	}()

	fmt.Printf("外部：a = %d, b = %s\n", a, str)
}
