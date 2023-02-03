/*
 * @Author MierX
 * @Title 03字符串处理函数.go
 * @Date 2023-02-03 周一
 * @Time 17:42:30
 */

package main

import (
	"fmt"
	"strings"
)

func main0301() {
	// strings 是字符串工具包
	// strings.Contains函数 查找一个字符串在另外一个字符串中是否存在 第一个形参是被查找的字符串 第二个形参是查找的字符串
	// 该函数采用的是 模糊查找
	str1 := "hello world!"
	str2 := "llo"
	rs := strings.Contains(str1, str2)
	fmt.Println(rs)
}

func main0302() {
	slice := []string{"你", "好", "世", "界", "！"}

	// 将切片里的元素用指定字符串连接起来返回一个字符串 第一个形参传入切片变量 第二个形参传入连接字符串
	str := strings.Join(slice, "")
	fmt.Println(str)
}

func main0303() {
	str1 := "hello world!"
	str2 := "llo"

	// 查找一个字符串在另一个字符串中首次匹配出现的下标，如果找不到则返回false 第一个形参传入被匹配的字符串 第二个形参传入要匹配的字符串
	i := strings.Index(str1, str2)
	fmt.Println(i)
}

func main0304() {
	str1 := "hello world!"

	// 将一个字符串按照指定次数重复拼接 第一个形参传入要重复的字符串 第二个形参传入要重复拼接的次数
	str2 := strings.Repeat(str1, 3)
	fmt.Println(str2)
}

func main() {
	str1 := "hello world! world"
	str2 := "china"

	// 将一个字符串中指定的部分进行替换，若该字符串中被替换的字符串有多处相同，则根据指定次数依次替换
	// 第一个形参传入被替换的字符串 第二个形参传入指定被替换的内容 第三个形参传入替换后的字符串 第四个形参传入替换次数
	// 如果替换次数小于0 则表示将所有匹配的部分都替换
	str3 := strings.Replace(str1, "world", str2, 1)
	fmt.Println(str3)
}
