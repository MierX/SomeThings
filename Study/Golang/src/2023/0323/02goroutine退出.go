/*
 * @Author MierX
 * @Title 02goroutine退出
 * @Date 2023-03-28 周一
 * @Time 09:47:25
 */

package main

import (
	"fmt"
	"time"
)

// 主协程退出了，其它子协程也会跟着退出
func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(time.Second)

		if i == 2 {
			break
		}
	}
}
