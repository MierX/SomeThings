/*
 * @Author MierX
 * @Title 07算术运算符.go
 * @Date 2023-01-09 周一
 * @Time 11:53:28
 */

package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	// 在除法计算时，除数不能为0
	// int类型数据运算，结果也是int类型
	fmt.Println(a / b)

	// 求余数（取模）
	fmt.Println(a % b)

	// 自增
	a++
	fmt.Println(a)

	// 自减
	a--
	fmt.Println(a)

	// 自增自减不能出现在表达式中

	// 编译：go build go文件
	// 编译并运行：go run go文件

	/*
	 * 异常分为三种：
	 * 1、编辑时异常
	 * 2、编译时异常
	 * 3、运行时异常
	 */
}
