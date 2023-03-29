/*
 * @Author MierX
 * @Title 03主协程先退出导致子协程没有来得及调用
 * @Date 2023-03-28 周一
 * @Time 09:56:50
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}
	}()
}
