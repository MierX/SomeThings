/*
 * @Author MierX
 * @Title 05map定义喝使用
 * @Date 2023-02-05 周日
 * @Time 14:08:59
 */

package main

import "fmt"

func main0501() {
	// map类型 键值对 key => value
	// map的定义 var 变量名 map[key类型]value类型
	var m map[int]string // 空变量
	fmt.Printf("%p\n", m)

	// map存储的方式，不是按照顺序存储的
	var m1 = map[int]string{101: "法师", 251: "张超", 666: "你好"}
	fmt.Printf("%p\n", m1)
	fmt.Println(m1)

	// 遍历map
	for i, s := range m1 {
		fmt.Println(i, s)
	}
}

func main() {
	m := make(map[string]int)
	fmt.Printf("%p\n", m)
	fmt.Println(len(m))
	m["key"] = 1
	fmt.Println(m)
	m["s"] = 1
	fmt.Println(m)
	m["s"] = 2
	fmt.Println(m)
}
