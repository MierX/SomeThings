/*
 * @Author MierX
 * @Title 15复数类型
 * @Date 2023-02-27 周一
 * @Time 17:00:34
 */

package main

import "fmt"

func main() {
	var t complex128 //声明
	t = 2.1 + 3.14i  //赋值
	fmt.Println("t = ", t)

	//自动推导类型
	t2 := 3.3 + 4.4i
	fmt.Printf("t2 类型是 %T\n", t2)

	//通过内建函数 取实部和虚部
	fmt.Println("real(t2) = ", real(t2), ", imag(t2) = ", imag(t2))
}
