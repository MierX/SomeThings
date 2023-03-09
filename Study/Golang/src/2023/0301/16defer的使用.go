/*
 * @Author MierX
 * @Title 16defer的使用
 * @Date 2023-03-09 周一
 * @Time 15:25:45
 */

package main

import "fmt"

func MyFunc25(x int) {
	rs := 100 / x
	fmt.Println("rs = ", rs)
}

func main() {
	fmt.Println("aaaaaaaaaaaaaaa")

	//defer 延迟调用 main函数结束前调用
	defer fmt.Println("bbbbbbbbbbb")
	//多个defer的执行顺序是先进后出
	defer MyFunc25(1)

	fmt.Println("ccccccccccc")
}
