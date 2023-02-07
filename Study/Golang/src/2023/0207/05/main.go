/*
 * @Author MierX
 * @Title main
 * @Date 2023-02-07 周一
 * @Time 17:36:25
 */

package main

// inter 定义接口
type inter interface {
	// CSocketProtocol 定义通信方法
	CSocketProtocol()

	// CEncDesProtocol 定义加密接口
	CEncDesProtocol()
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
