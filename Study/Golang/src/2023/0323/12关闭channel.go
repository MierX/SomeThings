/*
 * @Author MierX
 * @Title 12关闭channel
 * @Date 2023-03-28 周一
 * @Time 16:05:52
 */

package main

import "fmt"

func main() {
	ch := make(chan int)

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//关闭channel
		close(ch)
	}()

	for {
		//如果ok为true，说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else {
			//管道关闭
			break
		}
	}
}
