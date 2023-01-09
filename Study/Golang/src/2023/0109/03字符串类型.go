/**
 * @Author MierX
 * @Title 03字符串类型
 * @Date 2023-01-08 周日
 * @Time 12:43:28
 **/

package main

import "fmt"

func main() {
	// 双引号括起来的是字符串，单引号括起来的是字符
	var a string = "你好"

	fmt.Println(a)

	b := 'a'
	c := "a"
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	// len 是一个系统函数 用来计算字符串中字符的个数
	d := "hello world"
	fmt.Println(len(d))

	// 一个汉字占三个字节
	e := "你好"
	fmt.Println(len(e))

	f := d + e
	fmt.Println(f)
}
