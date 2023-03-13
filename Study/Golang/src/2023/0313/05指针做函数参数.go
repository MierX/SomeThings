/*
 * @Author MierX
 * @Title 05指针做函数参数
 * @Date 2023-03-13 周一
 * @Time 15:31:06
 */

package main

import "fmt"

func swap1(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("swap1:a = %d, b = %d\n", *a, *b)
}

func main() {
	a, b := 10, 20

	swap1(&a, &b)
	fmt.Printf("main:a = %d, b = %d\n", a, b)
}
