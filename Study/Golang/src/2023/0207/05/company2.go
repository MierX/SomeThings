/*
 * @Author MierX
 * @Title company2
 * @Date 2023-02-07 周一
 * @Time 17:39:34
 */

package main

import "fmt"

// CSckImp2 定义厂商2结构体
type CSckImp2 struct {
	data   string
	socket string
}

// CSocketProtocol 定义厂商2的通信方法
func (c2 *CSckImp2) CSocketProtocol() {
	fmt.Printf("厂商2的通信接口数据为：%s\n", c2.socket)
}

// CEncDesProtocol 定义厂商2的加密方法
func (c2 *CSckImp2) CEncDesProtocol() {
	fmt.Printf("厂商2的加密接口数据为：%s\n", c2.data)
}
