/*
 * @Author MierX
 * @Title 16格式化输出
 * @Date 2023-02-27 周一
 * @Time 17:07:56
 */

package main

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14
	//%T操作变量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)
	//%d 整型格式
	//%s 字符串格式
	//%c 字符格式
	//%f 浮点型格式
	fmt.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)
	//%v自动匹配格式输出
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)
}
