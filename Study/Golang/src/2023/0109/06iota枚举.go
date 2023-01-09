/**
 * @Author MierX
 * @Title 06iota枚举
 * @Date 2023-01-08 周日
 * @Time 14:23:33
 **/

package main

import "fmt"

func main() {
	// iota枚举 枚举格式如果写在一行中 值相等 如果换行 值递增
	const (
		a    = iota
		b, c = iota, iota
		d    = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// iota枚举 会自动给未赋值的变量按上一行递增
	const (
		e = iota
		f
		g
		h
	)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	// 定义iota枚举时，可以给变量自定义赋值，同时会给未赋值的变量与上一行相等，iota会根据当前变量行数赋值
	const (
		i = iota
		j = 100
		k
		l = iota
	)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
