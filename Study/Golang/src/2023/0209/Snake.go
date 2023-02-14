/*
 * @Author MierX
 * @Title Snake
 * @Date 2023-02-13 周一
 * @Time 16:49:18
 */

package main

import (
	"2023/0209/Clib"
	"time"
)

// Snake 定义蛇结构体
type Snake struct {
	size int        // 长度
	dir  int        // 方向
	pos  []Position // 坐标切片
}

// SnakeInit 初始化蛇信息
func (s *Snake) SnakeInit() {
	// 初始化地图
	MapInit()

	// 创建食物
	RandomFood(s)

	// 初始化蛇的长度
	s.size = 2

	// 设置蛇头位置
	s.pos = append(s.pos, Position{x: WIDE / 2, y: HIGH / 2}, Position{x: WIDE/2 - 1, y: HIGH / 2})

	// 用“wasd”控制上下左右(UDLR)
	s.dir = 'R'

	// 绘制蛇
	for i := 0; i < s.size; i++ {
		var ch = byte('*')
		if i == 0 {
			ch = '@'
		}
		ShowUI(s.pos[i].x, s.pos[i].y, ch)
	}

	// go 添加一个独立的函数 非阻塞执行
	// 接受键盘中的输入
	go func() {
		for {
			switch Clib.Direction() {
			// 根据键盘输入设置方向
			case 72, 87, 119:
				if s.dir != 'P' {
					if s.dir != 'D' {
						s.dir = 'U'
					}
				}
			case 80, 83, 115:
				if s.dir != 'P' {
					if s.dir != 'U' {
						s.dir = 'D'
					}
				}
			case 65, 75, 97:
				if s.dir != 'P' {
					if s.dir != 'R' {
						s.dir = 'L'
					}
				}
			case 68, 77, 100:
				if s.dir != 'P' {
					if s.dir != 'L' {
						s.dir = 'R'
					}
				}
			case 32:
				if s.dir != 'P' {
					lastDir = s.dir
					s.dir = 'P'
				} else {
					s.dir = lastDir
				}
			}
		}
	}()
}

// PlayGame 蛇流程控制
func (s *Snake) PlayGame() {
	// 蛇坐标偏移值
	var dx, dy int

	// 游戏流程控制
OuterLoop:
	for {
	FLAG:
		// 延迟执行 333秒
		time.Sleep(time.Duration(float64(time.Second) / level))

		// 判断暂停
		if s.dir == 'P' {
			goto FLAG
		}

		// 更新蛇的位置
		switch s.dir {
		case 'U':
			dx = 0
			dy = -1
		case 'D':
			dx = 0
			dy = 1
		case 'L':
			dx = -1
			dy = 0
		case 'R':
			dx = 1
			dy = 0
		}

		// 记录蛇末尾坐标
		lx := s.pos[s.size-1].x
		ly := s.pos[s.size-1].y

		// 蛇身移动
		for i := s.size - 1; i > 0; i-- {
			s.pos[i].x = s.pos[i-1].x
			s.pos[i].y = s.pos[i-1].y
		}

		// 蛇头移动
		s.pos[0].x += dx
		s.pos[0].y += dy

		// 绘制蛇
		for i := 0; i < s.size; i++ {
			var ch = byte('*')
			if i == 0 {
				ch = '@'
			}
			ShowUI(s.pos[i].x, s.pos[i].y, ch)
		}

		// 清除旧末尾
		ShowUI(lx, ly, ' ')

		// 蛇头和墙的碰撞判断
		if s.pos[0].x < 1 || s.pos[0].x >= WIDE || s.pos[0].y < 0 || s.pos[0].y >= HIGH-1 {
			break
		}

		// 蛇头和身体的碰撞判断
		for i := 1; i < s.size; i++ {
			if s.pos[0].x == s.pos[i].x && s.pos[0].y == s.pos[i].y {
				break OuterLoop
			}
		}

		// 蛇头和食物的碰撞判断
		if s.pos[0].x == food.x && s.pos[0].y == food.y {
			// 加长身体
			s.size++
			// 身体移动
			for i := s.size - 1; i > 0; i-- {
				if i == s.size-1 {
					s.pos = append(s.pos, Position{x: s.pos[i-1].x, y: s.pos[i-1].y})
				} else {
					s.pos[i].x = s.pos[i-1].x
					s.pos[i].y = s.pos[i-1].y
				}
			}
			s.pos[0].x += dx
			s.pos[0].y += dy

			// 随机新食物
			RandomFood(s)

			// 份数增加
			score++

			// 提升级别
			level += 0.01
		}
	}
}
