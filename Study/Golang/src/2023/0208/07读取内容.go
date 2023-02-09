/*
 * @Author MierX
 * @Title 07读取内容
 * @Date 2023-02-09 周一
 * @Time 11:56:56
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main0701() {
	// 块读取
	// 文件块读取（根据接收文件内容的切片字节变量的空间，一块一块的读取）

	file, err := os.OpenFile("./Study/Golang/src/2023/0208/test.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var buf []byte
	// 为字节切片创建空间
	buf = make([]byte, 1024)

	// 读取文件，将文件内容赋值给字节切片（需要有足够的空间）
	n, _ := file.Read(buf)
	fmt.Println(n)
	fmt.Println(string(buf[:n]))

	//buf1 := make([]byte, 3)
	//for {
	//	n1, err1 := file.Read(buf1)
	//	// io.EOF 表示的是文件的结尾 当Read读取到结尾时，返回的是 errors.New("EOF")
	//	if err1 == io.EOF {
	//		fmt.Println(err1)
	//		break
	//	}
	//
	//	fmt.Println(string(buf1[:n1]))
	//}
}

func main() {
	// 行读取
	fp, err := os.OpenFile("./Study/Golang/src/2023/0208/test.txt", os.O_RDONLY, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	// 创建文件缓冲区
	r := bufio.NewReader(fp)

	// 定义内容切割的标志字符串
	slice, _ := r.ReadBytes('\n')
	fmt.Println(string(slice))

	// 读取下一行内容
	slice, _ = r.ReadBytes('\n')
	fmt.Println(string(slice))
}
