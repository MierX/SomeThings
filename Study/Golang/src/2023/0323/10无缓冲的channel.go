/*
 * @Author MierX
 * @Title 10无缓冲的channel
 * @Date 2023-03-28 周一
 * @Time 15:39:27
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个无缓冲的channel
	ch := make(chan int, 0)

	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程 i = ", i)
			ch <- i //往通道里写入内容
		}
	}()

	//延时2s
	time.Sleep(time.Second * 2)
	for i := 0; i < 3; i++ {
		num := <-ch //读取通道中的内容，没有内容会阻塞
		fmt.Println("主协程 num = ", num)
	}
}
