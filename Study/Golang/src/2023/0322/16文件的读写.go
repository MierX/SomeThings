/*
 * @Author MierX
 * @Title 16文件的读写
 * @Date 2023-03-27 周一
 * @Time 16:06:39
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//创建文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer file.Close()

	//使用文件
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i) // 将 i = 0\n 这个字符串存储在buf中
		file.WriteString(buf)
	}
}

func ReadFile(path string) {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err1 = ", err)
		return
	}

	//关闭文件
	defer file.Close()

	buf := make([]byte, 1024*2) //2kb大小

	//n代表从文件读取内容的长度
	n, err := file.Read(buf)
	if err != nil && err != io.EOF { // io.EOF 代表 文件结束符
		fmt.Println("err2 = ", err)
		return
	}

	fmt.Println("buf:")
	fmt.Println(string(buf[:n]))
}

func ReadFileLine(path string) {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err1 = ", err)
		return
	}

	//关闭文件
	defer file.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(file)

	for {
		//以 \n 为结束读取标志
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				//文件已经读完，结束循环
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}
}

func main() {
	path := "Study/Golang/src/2023/0322/demo.txt"
	WriteFile(path)
	//ReadFile(path)

	//每次只读取一行
	ReadFileLine(path)
}
