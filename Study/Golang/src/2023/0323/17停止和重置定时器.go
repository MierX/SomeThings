/*
 * @Author MierX
 * @Title 17停止和重置定时器
 * @Date 2023-03-28 周一
 * @Time 16:41:00
 */

package main

import (
	"fmt"
	"time"
)

// 重置定时器
func main() {
	timer := time.NewTimer(10 * time.Second)
	timer.Reset(1 * time.Second) //重置为1s

	<-timer.C
	fmt.Println("子协程可以打印了，因为定时器的时间到")
}

// 停止定时器
func main01() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间到")
	}()

	timer.Stop() //停止定时器

	for {

	}
}
