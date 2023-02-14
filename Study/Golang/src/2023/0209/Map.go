/*
 * @Author MierX
 * @Title Map
 * @Date 2023-02-13 周一
 * @Time 16:48:30
 */

package main

import (
	"2023/0209/Clib"
	"fmt"
	"math/rand"
	"os"
	"time"
)

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

// Food 定义食物结构体
type Food struct {
	Position
}

// RandomFood 定义生成食物方法
func RandomFood(s *Snake) {
	food.x = rand.Intn(WIDE) // 0-19
	food.y = rand.Intn(HIGH)

	for i := 0; i < s.size-1; i++ {
		if food.x == s.pos[i].x && food.y == s.pos[i].y {
			food.x = rand.Intn(WIDE) // 0-19
			food.y = rand.Intn(HIGH)
			i = 0
		}
	}

	// 显示食物位置
	ShowUI(food.x, food.y, '#')
}

func PrintScore() {
	// 打印分数
	Clib.GotoPosition(0, 23)
	fmt.Fprintln(os.Stderr, "分数：", score)
	time.Sleep(time.Second * 2)
}
