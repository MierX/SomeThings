/*
 * @Author MierX
 * @Title 22切片做函数参数
 * @Date 2023-03-15 周一
 * @Time 16:50:13
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	n := 10

	//创建一个切片，len为n
	s := make([]int, n)

	InitData(s) //初始化数组
	fmt.Println("排序前：", s)

	BubbleSort(s) //冒泡怕徐
	fmt.Println("排序后：", s)
}
