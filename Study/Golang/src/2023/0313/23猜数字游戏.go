/*
 * @Author MierX
 * @Title 23猜数字游戏
 * @Date 2023-03-15 周一
 * @Time 16:57:15
 */

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func CreateNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	for {
		*p = rand.Intn(10000)
		if *p >= 1000 {
			break
		}
	}
}

func GetNum(s []int, num int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func OneGame(s []int) (int, int, string) {
	var num int
	count := 0
	keySlice := make([]int, 4)
	for {
		count++
		rs := true
		for {
			fmt.Printf("请输入一个4位数：")
			fmt.Scan(&num)
			if num <= 999 || num >= 10000 {
				fmt.Println("输入的数不符合要求")
				continue
			}
			break
		}

		GetNum(keySlice, num)

		for i := 0; i < len(s); i++ {
			if keySlice[i] > s[i] {
				rs = false
				fmt.Println("第" + strconv.Itoa(i+1) + "位数字猜错了，应该再小一点~")
				break
			} else if keySlice[i] < s[i] {
				rs = false
				fmt.Println("第" + strconv.Itoa(i+1) + "位数字猜错了，应该再大一点~")
				break
			}
		}

		if rs == true {
			break
		}
	}

	return num, count, "都猜对了！"
}

func main() {
	var randNum int

	//产生一个4位的随机数
	CreateNum(&randNum)

	randSlice := make([]int, 4)
	//保存这个4位数的每一位
	GetNum(randSlice, randNum)

	scan, count, result := OneGame(randSlice) //猜数字游戏
	fmt.Printf("题目：%d， 输入：%d，尝试次数：%d，结果：%v\n", randNum, scan, count, result)
}
