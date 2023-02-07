/*
 * @Author MierX
 * @Title 08方法的继承
 * @Date 2023-02-07 周一
 * @Time 15:33:36
 */

package main

import "fmt"

type person7 struct {
	id int
}

type student7 struct {
	person7
	class int
}

func (p *person7) Print() {
	fmt.Printf("编号：%d\n", p.id)
}

func main() {
	p := person7{101}
	p.Print()

	s := student7{person7{101}, 202}
	// 子类可以调用父类的方法 因为子类继承了父类
	// 父类不能调用子类的方法
	s.Print()
}
