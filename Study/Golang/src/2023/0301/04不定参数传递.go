/*
 * @Author MierX
 * @Title 04不定参数传递
 * @Date 2023-03-01 周一
 * @Time 16:32:52
 */

package main

import "fmt"

func MyFunc07(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func MyFunc08(args ...int) {
	MyFunc07(args...)

	//只传递前2个参数
	MyFunc07(args[:2]...)

	//不传递前两个参数
	MyFunc07(args[2:]...)
}

func main() {
	MyFunc08(1, 2, 3, 4, 5)
}
