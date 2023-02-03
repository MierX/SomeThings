/*
 * @Author MierX
 * @Title 05随机数练习.go
 * @Date 2023-02-02 周一
 * @Time 17:52:19
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main0501() {
	rand.Seed(time.Now().UnixNano())
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println(arr)
}

func main0502() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1

	value := 0
	for {
		fmt.Println("请输出一个数字")
		fmt.Scan(&value)

		if num > value {
			fmt.Println("您输入的数字太小了！")
		} else if num < value {
			fmt.Println("你输入的数字太大了！")
		} else {
			fmt.Println("你猜对了！")
			break
		}
	}
}

func main() {
	// 随机双色球彩票
	// 红色 1-33 选择6个 不能重复 蓝球 1-16 选择1个 可以和红球重复

	rand.Seed(time.Now().UnixNano())
	var redArr [6]int
	for i := 0; i < len(redArr); i++ {
		redArr[i] = rand.Intn(33) + 1
		for j := 0; j < i; j++ {
			// 数据重复
			if redArr[i] == redArr[j] {
				redArr[i] = rand.Intn(33) + 1

				// 将j赋值为-1是因为循环结束时，j会+1
				j = -1
			}
		}
	}
	blue := rand.Intn(16) + 1
	fmt.Println(redArr)
	fmt.Println(blue)
}
