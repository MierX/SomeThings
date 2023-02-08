/*
 * @Author MierX
 * @Title 08类型断言
 * @Date 2023-02-08 周一
 * @Time 13:55:51
 */

package main

import "fmt"

func main0801() {
	var i interface{}
	i = 10

	// 类型断言： 接收变量的值（根据判断类型定义该变量的类型），值的判断 := 接口变量.（要判断的数据类型）
	value, ok := i.(int)
	if ok {
		fmt.Println("整形数据：", value+10)
	} else {
		fmt.Println("错误")
	}
}

func demo() {
	fmt.Println("hello world")
}

func main() {
	var i []interface{}
	i = append(i, 10, 3.14, "hello world", demo)
	for _, i3 := range i {
		value, ok := i3.(int)
		if ok {
			fmt.Println("整形数据：", value+10)
		} else if value, ok := i3.(float64); ok {
			fmt.Println("浮点型数据：", value+10.0)
		} else if value, ok := i3.(string); ok {
			fmt.Println("字符串数据：", value+"!")
		} else if value, ok := i3.(func()); ok {
			value()
		} else {
			fmt.Println("错误")
		}
	}
}
