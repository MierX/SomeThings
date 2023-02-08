/*
 * @Author MierX
 * @Title 05新建文件
 * @Date 2023-02-08 周一
 * @Time 17:36:06
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件 os.Create(文件名【可以写绝对路径 也可以写相对路径（以文件运行的目录）】)
	// rs 是文件创建成功后的文件指针（内存地址） err 是文件创建失败后的报错信息
	// 如果文件不存在，会创建新文件 如果文件存在，会清空源文件内容
	rs, err := os.Create("./Study/Golang/src/2023/0208/test.txt")
	if err != nil {
		// 文件创建失败的场景：
		// 1、路径不存在
		// 2、文件权限不足
		// 3、程序打开文件上限 65535
		fmt.Println(err)
	} else {
		// 延迟调用关闭文件 主要是防止在程序最后忘记关闭文件
		defer rs.Close()

		// 文件创建成功可以操作文件
		fmt.Println(rs)
		fmt.Println(*rs)

		// 关闭文件
		// 如果打开文件不关闭 会造成内存浪费 占用程序打开文件的上限
		rs.Close()
	}
}
