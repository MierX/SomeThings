/*
 * @Author MierX
 * @Title 08通过channel实现同步
 * @Date 2023-03-28 周一
 * @Time 11:43:14
 */

package main

import (
	"fmt"
	"time"
)

// 全局变量，创建一个channel
var ch = make(chan int)

// Printer 定义一个打印机，参数为字符串，按每个字符打印
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

// Person01 Person01执行完才能执行Person02
func Person01() {
	Printer("hello")
	ch <- 666 //给管道写数据
}

func Person02() {
	<-ch //从管道取数据，接收，如果通道没有数据，他就会阻塞
	Printer("world")
}

func main() {
	//新建两个协程，两个协程同时使用打印机
	go Person01()
	go Person02()

	//特地不让主协程结束，死循环
	for {

	}
}
