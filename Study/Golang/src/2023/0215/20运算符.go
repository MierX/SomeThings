/*
 * @Author MierX
 * @Title 20运算符
 * @Date 2023-03-01 周一
 * @Time 15:03:35
 */

package main

import "fmt"

func main() {
	//关系运算符
	fmt.Println("4 > 3 结果：", 4 > 3)
	fmt.Println("4 != 3 结果：", 4 != 3)

	//逻辑运算符
	fmt.Println("!(4 > 3) 结果：", !(4 > 3))
	fmt.Println("!(4 != 3) 结果：", !(4 != 3))

	//&& 与，并且，两侧为真则真
	fmt.Println("true && true 结果：", true && true)
	fmt.Println("true && false 结果：", true && false)

	//|| 或，一侧为真则真
	fmt.Println("true || false 结果：", true || false)
	fmt.Println("false || false 结果：", false && false)

	a := 10
	fmt.Println("0 <= a && a <= 10 结果：", 0 <= a && a <= 10)
}
