/*
 * @Author MierX
 * @Title 29goto的使用
 * @Date 2023-03-01 周一
 * @Time 16:02:33
 */

package main

import "fmt"

func main() {
	//goto可以用在任何地方，但是不能跨函数使用
	fmt.Println("123")

	goto End //goto是关键字 End是用户定义的

	fmt.Println("456")

End:
	fmt.Println("789")
}
