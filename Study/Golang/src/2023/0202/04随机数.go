/*
 * @Author MierX
 * @Title 04随机数.go
 * @Date 2023-02-02 周一
 * @Time 17:44:34
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成随机数：
	// 1.导入头文件 math/rand time
	// 2.设置随机数种子
	// 3.创建随机数

	// 设置种子 进行数据混淆
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(10)) // 0-9
}
