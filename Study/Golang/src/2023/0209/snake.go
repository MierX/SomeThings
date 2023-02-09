/*
 * @Author MierX
 * @Title main
 * @Date 2023-02-09 周一
 * @Time 17:17:23
 */

package main

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
	size int        // 长度
	dir  int        // 方向
	pos  []Position // 坐标切片
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
}

// RandomFood 定义生成食物方法
func RandomFood() {

}

func main() {

}
