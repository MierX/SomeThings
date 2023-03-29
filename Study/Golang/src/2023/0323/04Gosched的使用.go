/*
 * @Author MierX
 * @Title 04Gosched的使用
 * @Date 2023-03-28 周一
 * @Time 10:02:40
 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片，先让别的协程执行，它执行完再回来执行此协程
		runtime.Gosched()
		fmt.Println("hello")
	}
}
