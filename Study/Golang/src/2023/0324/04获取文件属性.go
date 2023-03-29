/*
 * @Author MierX
 * @Title 04获取文件属性
 * @Date 2023-03-29 周一
 * @Time 14:54:19
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	//获取命令行输入内容
	list := os.Args
	if len(list) != 2 {
		fmt.Println("usage:xxx file")
		return
	}

	//获取文件名
	fileName := list[1]

	//获取文件属性
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("name = ", info.Name())
	fmt.Println("size = ", info.Size())
}
