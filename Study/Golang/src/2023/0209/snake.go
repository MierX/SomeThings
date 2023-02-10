/*
 * @Author MierX
 * @Title main
 * @Date 2023-02-09 周一
 * @Time 17:17:23
 */

package main

import (
	"2023/0209/Clib"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// WIDE 定义全局常量 宽度比例
const WIDE = 20

// HIGH 定义全局常量 长度比例
const HIGH = 20

// 定义全局变量 食物
var food Food

// Position 定义父类坐标结构体
type Position struct {
	x int // 横轴
	y int // 竖轴
}

// Snake 定义蛇结构体
type Snake struct {
	size int                   // 长度
	dir  int                   // 方向
	pos  [WIDE * HIGH]Position // 坐标切片
}

// Food 定义食物结构体
type Food struct {
	Position
}

// SnakeInit 初始化蛇信息
func (s *Snake) SnakeInit() {
	// 初始化蛇的长度
	s.size = 2

	// 设置蛇头位置
	s.pos[0].x = WIDE / 2
	s.pos[0].y = HIGH / 2

	// 设置蛇尾位置
	s.pos[1].x = WIDE/2 - 1
	s.pos[1].y = HIGH / 2

	// 用“WASD”控制上左下右
	s.dir = 'D'

	// 绘制蛇
	for i := 0; i < s.size; i++ {
		var ch = byte('*')
		if i == 0 {
			ch = '@'
		}
		ShowUI(s.pos[i].x, s.pos[i].y, ch)
	}
}

func (s *Snake) PlayGame() {
	// 游戏流程控制
	for {
		// 蛇头和墙的碰撞判断
		if s.pos[0].x < 0 || s.pos[0].x >= WIDE || s.pos[0].y < 0 || s.pos[0].y >= HIGH {
			return
		}

		// 蛇头和身体的碰撞判断
		for i := 1; i < s.size; i++ {
			if s.pos[0].x == s.pos[i].x || s.pos[0].y == s.pos[i].y {
				return
			}
		}

		// 蛇头和食物的碰撞判断
		if s.pos[0].x == food.x && s.pos[0].y == food.y {
			// 加长身体
			s.size++

			// 随机新食物
			RandomFood()
		}
	}
}

// RandomFood 定义生成食物方法
func RandomFood() {
	food.x = rand.Intn(WIDE) // 0-19
	food.y = rand.Intn(HIGH)

	//TODO 与蛇重合的话，重新随机

	// 显示食物位置
	ShowUI(food.x, food.y, '#')
}

// MapInit 设置地图
func MapInit() {
	// 输出初始画面
	fmt.Fprintln(os.Stderr,
		`
  #--------------------------------------#
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  |                                      |
  #--------------------------------------#
`)
}

// ShowUI 绘制元素
func ShowUI(x, y int, ch byte) {
	// 调用C语言代码 设置控制台的光标
	// 根据地图坐标有偏移
	Clib.GotoPosition(x*2+2, y+2)
	fmt.Fprintf(os.Stderr, "%c", ch)
}

func main() {
	// 设置随机数种子 用作于混淆
	rand.Seed(time.Now().UnixNano())

	// 隐藏控制台的光标
	Clib.HideCursor()

	// 初始化地图
	MapInit()

	// 创建食物
	RandomFood()

	// 创建蛇对象
	var snake Snake

	// 初始化蛇的信息
	snake.SnakeInit()
}
