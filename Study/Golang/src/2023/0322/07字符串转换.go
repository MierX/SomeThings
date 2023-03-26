/*
 * @Author MierX
 * @Title 07字符串转换
 * @Date 2023-03-26 周一
 * @Time 16:45:47
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个为要追加的数，第三个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abc")

	fmt.Println("slice = ", string(slice))

	//其它类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str = ", str)

	//整型转字符串
	str = strconv.Itoa(6666)
	fmt.Println("str = ", str)

	//字符串转其它类型
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	//把字符串转换为整型
	a, _ := strconv.Atoi("567")
	fmt.Println("a = ", a)
}
