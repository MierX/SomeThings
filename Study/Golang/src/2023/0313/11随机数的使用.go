/*
 * @Author MierX
 * @Title 11随机数的使用
 * @Date 2023-03-13 周一
 * @Time 16:06:53
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子
	//如果种子参数一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano()) //以当前系统时间作为种子参数

	for i := 0; i < 5; i++ {
		//产生随机数
		fmt.Println("rand = ", rand.Int())     //随机很大的数
		fmt.Println("rand = ", rand.Intn(100)) //限制在100内的数
	}
}
