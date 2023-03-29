/*
 * @Author MierX
 * @Title 14单向channel的应用
 * @Date 2023-03-28 周一
 * @Time 16:22:31
 */

package main

import "fmt"

// 此函数只能写入通道
func writeCh(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("i = ", i)
	}
	close(ch)
}

// 此函数只能读取通道
func readCh(ch <-chan int) {
	for num := range ch {
		fmt.Println("num = ", num)
	}
}

func main() {
	//创建一个双向通道
	ch := make(chan int)

	//写入通道
	go writeCh(ch) //channel传参，是引用传递

	//读取通道
	readCh(ch)
}
