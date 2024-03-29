/*
 * @Author MierX
 * @Title 14字符和字符串的区别
 * @Date 2023-02-27 周一
 * @Time 16:27:29
 */

package main

import "fmt"

func main() {
	var ch byte
	var str string

	//字符
	//1、单引号
	//2、字符，往往都只有一个字符，转义字符除外
	ch = 'a'
	fmt.Println(ch)

	//字符串
	//1、双引号
	//2、字符串有1个或多个字符组成
	//3、字符串都是隐藏了一个结束符，"\0"
	str = "a" //由'a'和'\0'组成了一个字符串
	fmt.Println("str = ", str)

	str = "hello go"
	//只想操作字符串的某个字符，从0开始操作
	fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])
}
