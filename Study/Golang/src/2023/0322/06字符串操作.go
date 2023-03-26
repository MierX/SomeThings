/*
 * @Author MierX
 * @Title 06字符串操作
 * @Date 2023-03-26 周一
 * @Time 16:28:46
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("hello go", "hello"))

	//Join组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "@")
	fmt.Println("buf = ", buf)

	//Index，查找子串的位置
	fmt.Println(strings.Index("abc hello", "hello"))

	//Repeat
	buf = strings.Repeat("go", 3)
	fmt.Println("buf = ", buf)

	//Split，以指定的分隔符拆分
	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2)

	//Trim去掉两头空格
	fmt.Println(strings.Trim("    are u ok?     ", " "))

	//去掉空格，把单词放入切片中
	fmt.Println(strings.Fields("    are u ok?     "))
}
