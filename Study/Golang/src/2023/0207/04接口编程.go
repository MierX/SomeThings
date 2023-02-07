/*
 * @Author MierX
 * @Title 04接口编程
 * @Date 2023-02-07 周一
 * @Time 17:21:12
 */

package main

import "fmt"

// inter 定义接口
type inter interface {
	// CSocketProtocol 定义通信方法
	CSocketProtocol()

	// CEncDesProtocol 定义加密接口
	CEncDesProtocol()
}

// CSckImp1 定义厂商1结构体
type CSckImp1 struct {
	data   string
	socket string
}

// CSckImp2 定义厂商2结构体
type CSckImp2 struct {
	data   string
	socket string
}

// CSocketProtocol 定义厂商1的通信方法
func (c1 *CSckImp1) CSocketProtocol() {
	fmt.Printf("厂商1的通信接口数据为：%s\n", c1.socket)
}

// CEncDesProtocol 定义厂商1的加密方法
func (c1 *CSckImp1) CEncDesProtocol() {
	fmt.Printf("厂商1的加密接口数据为：%s\n", c1.data)

}

// CSocketProtocol 定义厂商2的通信方法
func (c2 *CSckImp2) CSocketProtocol() {
	fmt.Printf("厂商2的通信接口数据为：%s\n", c2.socket)
}

// CEncDesProtocol 定义厂商2的加密方法
func (c2 *CSckImp2) CEncDesProtocol() {
	fmt.Printf("厂商2的加密接口数据为：%s\n", c2.data)
}

// 多态实现
func framework(i inter) {
	i.CSocketProtocol()
	i.CEncDesProtocol()
}

func main() {
	c1 := CSckImp1{"厂商1的加密数据", "厂商1的通信数据"}
	framework(&c1)
	c2 := CSckImp2{"厂商2的加密数据", "厂商2的通信数据"}
	framework(&c2)
}
