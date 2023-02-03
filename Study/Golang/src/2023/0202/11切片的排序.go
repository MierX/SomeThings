/*
 * @Author MierX
 * @Title 11切片的排序.go
 * @Date 2023-02-03 周一
 * @Time 16:14:43
 */

package main

import "fmt"

func main() {
	slice := []int{9, 1, 5, 6, 7, 4, 8, 3, 2, 0}
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println(slice)
}
