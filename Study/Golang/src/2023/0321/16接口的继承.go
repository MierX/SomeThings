/*
 * @Author MierX
 * @Title 16接口的继承
 * @Date 2023-03-24 周一
 * @Time 16:59:00
 */

package main

import "fmt"

// Humaner01 定义接口类型
type Humaner01 interface { //子集
	//方法，只有生命，没有实现，由别的类型（自定义类型）实现
	sayHi()
}

// Person13 定义接口类型
type Person13 interface { //超集
	Humaner01 //匿名字段，继承了sayHi()
	sing(lrc string)
}

type Student08 struct {
	name string
	id   int
}

// Student08实现了sayHi方法
func (tmp *Student08) sayHi() {
	fmt.Printf("Student07[%s, %d] sayHi\n", tmp.name, tmp.id)
}

func (tmp *Student08) sing(lrc string) {
	fmt.Println("Student08在唱着：", lrc)
}

func main() {
	//定义一个接口类型的变量
	var i Person13
	s := &Student08{"mike", 666}
	i = s
	i.sayHi() //继承过来的方法
	i.sing("hi")
}
