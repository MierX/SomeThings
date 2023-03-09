/*
 * @Author MierX
 * @Title 17defer和匿名函数结合使用
 * @Date 2023-03-09 周一
 * @Time 15:35:22
 */

package main

import "fmt"

func main01() {
	a := 10
	b := 20

	defer func() {
		fmt.Printf("内部：a = %d, b = %d\n", a, b)
	}()

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)
}

func main() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("内部：a = %d, b = %d\n", a, b)
	}(a, b)

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)
}
