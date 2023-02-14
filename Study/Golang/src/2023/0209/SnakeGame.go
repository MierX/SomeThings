/*
 * @Author MierX
 * @Title main
 * @Date 2023-02-09 周一
 * @Time 17:17:23
 */

package main

import (
	"2023/0209/Clib"
	"math/rand"
	"time"
)

// WIDE 定义全局常量 宽度比例
const WIDE = 20

// HIGH 定义全局常量 长度比例
const HIGH = 20

// 定义全局变量 食物
var food Food

// 定义分数
var score int

// 定义级别
var level = 1.00

// 定义上一个操作
var lastDir int = 'R'

// Position 定义父类坐标结构体
type Position struct {
	x int // 横轴
	y int // 竖轴
}

func main() {
	// 设置随机数种子 用作于混淆
	rand.Seed(time.Now().UnixNano())

	// 隐藏控制台的光标
	Clib.HideCursor()

	// 创建蛇对象
	var snake Snake

	// 初始化蛇的信息
	snake.SnakeInit()

	// 开始游戏
	snake.PlayGame()

	// 打印份数
	PrintScore()
}
