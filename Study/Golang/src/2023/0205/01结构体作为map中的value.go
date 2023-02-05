/*
 * @Author MierX
 * @Title 01结构体作为map中的value
 * @Date 2023-02-05 周日
 * @Time 15:58:03
 */

package main

import "fmt"

type student struct {
	name  string
	sex   string
	age   int
	score int
	addr  string
}

func main0101() {
	m := make(map[int]student)

	m[101] = student{"擎天柱", "男", 100, 100, "赛博坦"}
	fmt.Println(m)
	delete(m, 101)
	fmt.Println(m)
}

func main() {
	m := make(map[int][]student)

	m[101] = []student{
		{"擎天柱", "男", 100, 100, "赛博坦"},
		{"擎天柱", "男", 100, 100, "赛博坦"},
		{"擎天柱", "男", 100, 100, "赛博坦"},
	}
	fmt.Println(m)
	m[101] = append(m[101], student{"擎天柱", "男", 100, 100, "赛博坦"})
	fmt.Println(m)
}
