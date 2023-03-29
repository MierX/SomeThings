/*
 * @Author MierX
 * @Title 15Timer的使用
 * @Date 2023-03-28 周一
 * @Time 16:31:26
 */

package main

import (
	"fmt"
	"time"
)

// 验证time.NewTimer()时间到了只会相应一次
func main() {
	//创建一个定时器，设置时间为2s，2s后往time通道写内容（当前时间）
	timer := time.NewTimer(time.Second)

	for {
		<-timer.C //只会写一次
		fmt.Println("时间到")
	}
}

func main01() {
	//创建一个定时器，设置时间为2s，2s后往time通道写内容（当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	//2s后，往timer.C写入当前时间，有数据后就可以读取
	time := <-timer.C //没有数据会阻塞
	fmt.Println("time：", time)
}
