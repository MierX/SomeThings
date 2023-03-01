/*
 * @Author MierX
 * @Title 19类型别名
 * @Date 2023-03-01 周一
 * @Time 14:33:12
 */

package main

import "fmt"

func main() {
	//给int64起一个别名叫bigint
	type bigint int64
	var a bigint //等价于var a int64
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)

	var b long = 11
	fmt.Printf("b type is %T\n", b)
	var c char = 'a'
	fmt.Printf("c type is %T\n", c)
}
