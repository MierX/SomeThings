/*
 * @Author MierX
 * @Title 19斐波那契数列
 * @Date 2023-03-28 周一
 * @Time 17:05:05
 */

package main

import "fmt"

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//监听channel数据的流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}

func main() {
	ch := make(chan int)    //数字的通道
	quit := make(chan bool) //程序是否结束的通道

	//消费者：从channel中读取内容
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println("num = ", num)
		}

		//可以停止了
		quit <- true
	}()

	//生产者，往channel中写入数字
	fibonacci(ch, quit)
}
