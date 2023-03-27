/*
 * @Author MierX
 * @Title 15设备文件的使用
 * @Date 2023-03-27 周一
 * @Time 16:01:50
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Stdout.Close() //关闭标准设备文件，关闭后无法输出
	fmt.Println("are u ok?") //往标准输入设备（显示器）写内容

	//标准设备文件，默认打开，用户可以直接使用
	//os.Stdout
	os.Stdout.WriteString("are u ok?\n")

	//os.Stdout.Close() //关闭标准设备文件，关闭后无法输入
	var a int
	fmt.Println("请输入a：")
	fmt.Scan(&a) //从标准输入设备中读取内容放在a中
	fmt.Println("a = ", a)
}
