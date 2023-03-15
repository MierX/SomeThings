/*
 * @Author MierX
 * @Title 21copy的使用
 * @Date 2023-03-15 周一
 * @Time 16:46:21
 */

package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6, 6}
	copy(dstSlice, srcSlice)
	fmt.Println("dst = ", dstSlice)
	fmt.Println("src = ", srcSlice)
}
