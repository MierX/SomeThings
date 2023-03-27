/*
 * @Author MierX
 * @Title 17拷贝文件
 * @Date 2023-03-27 周一
 * @Time 16:28:50
 */

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件不能重名")
		return
	}

	//只读方式打开源文件
	sF, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("err1 = ", err)
		return
	}

	//新建目的文件
	dF, err := os.Create(dstFileName)
	if err != nil {
		fmt.Println("err2 = ", err)
		return
	}

	//关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理，从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 1024*4) //4kb大小缓冲区
	for {
		n, error := sF.Read(buf)
		if error != nil {
			if error == io.EOF {
				break
			}
			fmt.Println("error = ", error)
		}
		dF.Write(buf[:n])
	}
}
