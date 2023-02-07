/*
 * @Author MierX
 * @Title company1
 * @Date 2023-02-07 周一
 * @Time 17:39:26
 */

package main

import "fmt"

// CSckImp1 定义厂商1结构体
type CSckImp1 struct {
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
