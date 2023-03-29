/*
 * @Author MierX
 * @Title 09通过channel实现同步和数据交互
 * @Date 2023-03-28 周一
 * @Time 11:54:31
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("主协程也结束")

	//创建channel
	ch := make(chan string)

	go func() {
		defer fmt.Println("子协程调用完毕")

		for i := 0; i < 2; i++ {
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}

		ch <- "子协程工作完毕"
	}()

	str := <-ch //没有数据时阻塞
	fmt.Println("str = ", str)
}
