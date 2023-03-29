/*
 * @Author MierX
 * @Title 05Goexit的使用
 * @Date 2023-03-28 周一
 * @Time 10:06:28
 */

package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("c")

	//return
	runtime.Goexit() //终止所在的协程

	fmt.Println("d")
}

func main() {
	//创建子协程
	go func() {
		fmt.Println("a")

		//调用别的函数
		test()

		fmt.Println("b")
	}()

	//特地写一个死循环，目的不让主协程结束
	for {

	}
}
