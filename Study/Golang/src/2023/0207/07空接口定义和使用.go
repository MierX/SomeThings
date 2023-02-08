/*
 * @Author MierX
 * @Title 07空接口定义和使用
 * @Date 2023-02-08 周一
 * @Time 13:40:25
 */

package main

import "fmt"

func main0701() {
	// 定义一个空接口
	var i interface{}

	// 空接口的类型是nil
	fmt.Printf("%T\n", i) // <nil>

	// 接口变量可以接收任意类型的数据
	i = 10
	fmt.Println(i)
	fmt.Printf("%T\n", i) // int
	i = .14
	fmt.Println(i)
	fmt.Printf("%T\n", i) // float64
	i = "你好"
	fmt.Println(i)
	fmt.Printf("%T\n", i) // string

	// 接口类型 不能直接进行类型转换 需要使用类型断言
	// i = 10
	// a := 20
	// fmt.Println(i + a) 不能参与运算
}

func test() {
	fmt.Println("hello world")
}

func main() {
	// 定义一个 空接口类型的切片
	var i []interface{}
	fmt.Printf("%T\n", i)

	i = append(i, 10, 3.14, "hello world", test)

	for j := 0; j < len(i); j++ {
		fmt.Println(i[j])
	}
}
