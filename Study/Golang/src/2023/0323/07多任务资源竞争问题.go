/*
 * @Author MierX
 * @Title 07多任务资源竞争问题
 * @Date 2023-03-28 周一
 * @Time 10:28:10
 */

package main

import (
	"fmt"
	"time"
)

// Printer 定义一个打印机，参数为字符串，按每个字符打印
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func Person01() {
	Printer("hello")
}

func Person02() {
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
