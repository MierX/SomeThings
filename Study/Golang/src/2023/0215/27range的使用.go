/*
 * @Author MierX
 * @Title 27range的使用
 * @Date 2023-03-01 周一
 * @Time 15:52:20
 */

package main

import "fmt"

func main() {
	str := "abc"

	//通过for打印每个字符
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	//通过迭代打印每个元素，默认返回2个值，元素下标和元素本身
	for s, i2 := range str {
		fmt.Println(s, i2)
	}

	//通过匿名变量丢弃第二个返回值
	for i, _ := range str {
		fmt.Println(i)
	}
}
