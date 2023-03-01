/*
 * @Author MierX
 * @Title 05一个返回值
 * @Date 2023-03-01 周一
 * @Time 16:40:37
 */

package main

import "fmt"

// MyFunc09 无参有返回值 只有一个返回值
func MyFunc09() int { //有返回值的函数需要通过return中断函数 通过return返回
	return 666
}

// MyFunc10 给返回值起一个变量名 go语言推荐写法
func MyFunc10() (result int) { //有返回值的函数需要通过return中断函数 通过return返回
	result = 666
	return
}

func main() {
	fmt.Println(MyFunc09())
	fmt.Println(MyFunc10())
}
