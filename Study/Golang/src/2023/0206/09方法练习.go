/*
 * @Author MierX
 * @Title 09方法练习
 * @Date 2023-02-07 周一
 * @Time 15:41:35
 */

package main

import "fmt"

type human struct {
	name string
	age  int
	sex  string
}

type rep struct {
	human
	hobby string
}

type dev struct {
	human
	workTime int
}

func (h *human) SayHi() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我是%s生\n", h.name, h.age, h.sex)
}

func (r *rep) SayHello() {
	r.SayHi()
	fmt.Printf("我的爱好是%s\n", r.hobby)
}

func (d *dev) SayHello() {
	d.SayHi()
	fmt.Printf("我的工作年限是%d年\n", d.workTime)
}

func main() {
	rep := rep{human{"卓伟", 40, "男"}, "偷拍"}
	rep.SayHello()

	dev := dev{human{"图灵", 32, "男"}, 10}
	dev.SayHello()
}
