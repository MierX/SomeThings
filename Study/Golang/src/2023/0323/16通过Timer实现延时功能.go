/*
 * @Author MierX
 * @Title 16通过Timer实现延时功能
 * @Date 2023-03-28 周一
 * @Time 16:36:55
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	//延时2s后打印一句话
	<-time.After(2 * time.Second) //定时2s，阻塞2s，2s后会产生一个事件：往channel写内容
	fmt.Println("时间到")
}

func main02() {
	//延时2s后打印一句话
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
}

func main01() {
	//延时2s后打印一句话
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")
}
