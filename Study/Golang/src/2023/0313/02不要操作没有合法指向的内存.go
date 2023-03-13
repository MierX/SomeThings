/*
 * @Author MierX
 * @Title 02不要操作没有合法指向的内存
 * @Date 2023-03-13 周一
 * @Time 15:19:20
 */

package main

import "fmt"

func main() {
	var p *int
	p = nil
	fmt.Println("p = ", p)

	//*p = 666 //报错，因为p没有合法指向

	var a int
	p = &a
	*p = 666
	fmt.Println("a = ", a)
}
