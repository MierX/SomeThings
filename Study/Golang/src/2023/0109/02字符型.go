/**
 * @Author MierX
 * @Title 02字符型
 * @Date 2023-01-08 周日
 * @Time 12:24:55
 **/

package main

import "fmt"

func main() {
	// byte 类型相同与 uint8 类型
	var a byte = 'a'
	// 所有的字符都对应ASCII码值表中的整型数据
	fmt.Println(a)
	// %c 是一个占位符 表示打印输出一个字符
	fmt.Printf("%c\n", a)

	fmt.Printf("%c\n", 97)

	fmt.Printf("%T\n", a)

	fmt.Printf("%c\n", a-32)

	var b byte = '\n'
	fmt.Println(b)
}
