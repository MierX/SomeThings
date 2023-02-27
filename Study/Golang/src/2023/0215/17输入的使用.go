/*
 * @Author MierX
 * @Title 17输入的使用
 * @Date 2023-02-27 周一
 * @Time 17:16:25
 */

package main

import "fmt"

func main() {
	var a int //声明变量
	fmt.Printf("请输入变量a：")
	//阻塞等待用户的输入
	fmt.Scan(&a) //别忘了 &
	fmt.Println("a = ", a)
}
