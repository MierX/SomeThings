/*
 * @Author MierX
 * @Title 17接口转换
 * @Date 2023-03-24 周一
 * @Time 17:09:12
 */

package main

import "fmt"

// Humaner02 定义接口类型
type Humaner02 interface { //子集
	//方法，只有生命，没有实现，由别的类型（自定义类型）实现
	sayHi()
}

// Person14 定义接口类型
type Person14 interface { //超集
	Humaner02 //匿名字段，继承了sayHi()
	sing(lrc string)
}

type Student09 struct {
	name string
	id   int
}

// Student09实现了sayHi方法
func (tmp *Student09) sayHi() {
	fmt.Printf("Student07[%s, %d] sayHi\n", tmp.name, tmp.id)
}

func (tmp *Student09) sing(lrc string) {
	fmt.Println("Student09在唱着：", lrc)
}

func main() {
	//超集可以转换为子集，反过来不可以
	var iPro Person14 //超集
	iPro = &Student09{"mike", 666}
	var i Humaner02 //子集

	i = iPro
	i.sayHi()
}
