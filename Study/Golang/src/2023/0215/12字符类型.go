/*
 * @Author MierX
 * @Title 12字符类型
 * @Date 2023-02-27 周一
 * @Time 16:09:25
 */

package main

import "fmt"

func main() {
	var ch byte //声明字符类型
	ch = 97
	//格式化输出，%c以字符方式打印，%d以整型方式打印
	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'a' //字符，单引号
	fmt.Printf("%c, %d\n", ch, ch)

	//大写转小写，小写转大写
	fmt.Printf("大写：%d，小写：%d\n", 'A', 'a')
	fmt.Printf("大写转小写：%c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)

	//'\'以反斜杠开头的字符，是转义字符
	fmt.Printf("hello go%c", '\n')
	fmt.Printf("hello world")
}
