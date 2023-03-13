/*
 * @Author MierX
 * @Title 03new函数的使用
 * @Date 2023-03-13 周一
 * @Time 15:21:48
 */

package main

import "fmt"

func main() {
	//a := 10 //整型变量

	var p *int
	//指向一个合法内存
	//p = &a

	//p是*int，指向int类型
	p = new(int)
	fmt.Printf("p = %v\n", p)

	*p = 666
	fmt.Println("p = ", p)
}
