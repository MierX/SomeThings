/*
 * @Author MierX
 * @Title 18Ticker的使用
 * @Date 2023-03-28 周一
 * @Time 16:51:04
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		<-ticker.C

		i++
		fmt.Println("i = ", i)

		if i == 5 {
			ticker.Stop()
			break
		}

		if i > 2 {
			ticker.Reset(time.Duration(i) * time.Second)
		}
	}
}
