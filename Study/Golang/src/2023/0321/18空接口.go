/*
 * @Author MierX
 * @Title 18空接口
 * @Date 2023-03-24 周一
 * @Time 17:19:13
 */

package main

import "fmt"

func main() {
	//空接口万能类型，保存任意类型的
	var i interface{} = 1
	fmt.Println("i = ", i)

	i = "abc"
	fmt.Println("i = ", i)
}
