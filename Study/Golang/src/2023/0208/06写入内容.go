/*
 * @Author MierX
 * @Title 06写入内容
 * @Date 2023-02-09 周一
 * @Time 10:02:39
 */

package main

import (
	"fmt"
	"io"
	"os"
)

func main0601() {
	// WriteString 写入文件
	rs, err := os.Create("./Study/Golang/src/2023/0208/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 延迟关闭
	defer rs.Close()

	// 将字符串写入文件
	rs.WriteString("hello world\r\n")
	rs.WriteString("hello world!!!\r\n")

	s := []byte("hello\r\n")
	n, err := rs.Write(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}

func main0602() {
	// Write 写入文件
	rs, err := os.Create("./Study/Golang/src/2023/0208/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 延迟关闭
	defer rs.Close()

	// 将字符串写入文件
	s := []byte("hello\r\n")
	n, err1 := rs.Write(s)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(n)
}

func main0603() {
	// WriteAt 写入文件
	rs, err := os.Create("./Study/Golang/src/2023/0208/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 延迟关闭
	defer rs.Close()

	// 将字符串写入文件
	n, err1 := rs.WriteAt([]byte("你好\r\n"), 0)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(n)
}

func main() {
	// 打开文件

	// os.Open 打开文件 只读方式打开文件
	// rs, err := os.Open("./Study/Golang/src/2023/0208/test.txt")

	// os.OpenFile 根据权限值打开文件 第一个参数 文件路径 第二个参数 打开方式 第三个参数 控制文件模式（权限）
	rs, err := os.OpenFile("./Study/Golang/src/2023/0208/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rs.Close()

	// 获取文件的光标位置（通常是末尾）
	index, err2 := rs.Seek(0, io.SeekEnd)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	n, err1 := rs.WriteAt([]byte("hello\r\n"), index)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(n)
}
