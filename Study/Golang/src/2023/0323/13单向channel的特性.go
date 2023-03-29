/*
 * @Author MierX
 * @Title 13单向channel的特性
 * @Date 2023-03-28 周一
 * @Time 16:17:20
 */

package main

func main() {
	//创建一个channel，双向的
	ch := make(chan int)

	//双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch //只能写不能读
	var readCh <-chan int = ch  //只能读不能写

	writeCh <- 1
	<-readCh
}
