/*
 * @Author MierX
 * @Title 01创建goroutine
 * @Date 2023-03-27 周一
 * @Time 17:31:10
 */

package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask goroutine")
		time.Sleep(time.Second) //延时1秒
	}
}

func main() {
	go newTask() //新建一个协程

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second) //延时1秒
	}
}
