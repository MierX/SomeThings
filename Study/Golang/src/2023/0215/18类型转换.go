/*
 * @Author MierX
 * @Title 18类型转换
 * @Date 2023-03-01 周一
 * @Time 14:20:59
 */

package main

import "fmt"

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	//bool类型不能转换成int类型
	//fmt.Printf("flag = %d\n", int(flag))

	//0就是假，非0就是真
	//整型也不能转换为bool
	//flag = bool(1)

	var ch byte
	ch = 'a' //字符类型本质上就是整型
	var t int
	t = int(ch) //类型转换，把ch的值取出来后转成int，再给t赋值
	fmt.Println("t = ", t)
}
