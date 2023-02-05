/*
 * @Author MierX
 * @Title 04字符串类型转换
 * @Date 2023-02-05 周日
 * @Time 13:41:22
 */

package main

import (
	"fmt"
	"strconv"
)

func main0401() {
	str := "hello world!"
	// 将字符串转换成字符切片  强制类型转换
	slice := []byte(str)
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("%c", slice[i])
	}
}

func main0402() {
	slice := []byte{'h', 'e', 'l', 'l', 'o', 97}
	fmt.Println(string(slice))
}

func main0403() {
	// 将其他类型转换成字符串 使用 Format
	str := strconv.FormatBool(true)
	fmt.Println(str)
	fmt.Printf("%T\n", str)

	str1 := strconv.FormatInt(123, 2)
	fmt.Println(str1)
	fmt.Printf("%T\n", str1)

	str2 := strconv.FormatFloat(3.14, 'f', 6, 64)
	fmt.Println(str2)
	fmt.Printf("%T\n", str2)

	str4 := strconv.Itoa(123)
	fmt.Println(str4)
	fmt.Printf("%T\n", str4)
}

func main0404() {
	// 将字符串转成其它类型 第一个变量用于接收转换后的结果 第二个变量用于接收错误
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Printf("类型转化出错")
	} else {
		fmt.Println(b)
		fmt.Printf("%T\n", b)
	}

	i, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Printf("类型转化出错")
	} else {
		fmt.Println(i)
		fmt.Printf("%T\n", i)
	}

	f, err := strconv.ParseFloat("3.14159", 64)
	if err != nil {
		fmt.Printf("类型转化出错")
	} else {
		fmt.Println(f)
		fmt.Printf("%T\n", f)
	}
}

func main() {
	// 将其他类型转换成字符串添加到字符切片里面
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, false)
	slice = strconv.AppendInt(slice, 123, 2)
	slice = strconv.AppendFloat(slice, 3.14159, 'f', 4, 64)
	slice = strconv.AppendQuote(slice, "hello")
	fmt.Println(slice)
	fmt.Println(string(slice))
}
