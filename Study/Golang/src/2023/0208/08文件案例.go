/*
 * @Author MierX
 * @Title 08文件案例
 * @Date 2023-02-09 周一
 * @Time 15:52:36
 */

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 读取文件
	fp1, err1 := os.Open("./Study/Golang/src/2023/0208/test.jpeg")

	// 写入文件
	fp2, err2 := os.Create("./Study/Golang/src/2023/0208/test2.jpeg")

	if err1 != nil || err2 != nil {
		fmt.Println("操作文件失败")
		return
	}

	// 关闭文件
	defer fp1.Close()
	defer fp2.Close()

	// 拷贝文件
	// 通过 read 进行文件读取
	// 通过 write 进行写入
	buf := make([]byte, 1024)
	for {
		n, err3 := fp1.Read(buf)
		if err3 == io.EOF {
			fmt.Println("读取失败")
			break
		}

		fp2.Write(buf[:n])
	}
}
