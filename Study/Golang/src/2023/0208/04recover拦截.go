/*
 * @Author MierX
 * @Title 04recover拦截
 * @Date 2023-02-08 周一
 * @Time 16:58:39
 */

package main

import "fmt"

func demo(i int) {
	var arr [10]int

	// 错误拦截 一定要写在出现错误的代码之前
	defer func() {
		// 错误拦截 panic异常错误
		err := recover()
		fmt.Println(err)
	}()
	arr[i] = 123

	var p *int
	p = new(int)
	*p = 123

	fmt.Println(arr)
}

func main() {
	demo(1)
}
