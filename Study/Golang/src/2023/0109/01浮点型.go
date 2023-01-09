/**
 * @Author MierX
 * @Title 01浮点型
 * @Date 2023-01-08 周日
 * @Time 12:10:14
 **/

package main

import "fmt"

func main001801() {
	// 浮点型数据 分为 单精度浮点型float32（小数位数7位） 双精度浮点型float64（小数位数15位）
	// float64更精准
	var a float64 = 123.456

	fmt.Printf("%f\n", a)
	fmt.Printf("%.15f\n", a)
	fmt.Printf("%.20f\n", a)

	var b float32 = 123.456

	fmt.Printf("%f\n", b)
	fmt.Printf("%.7f\n", b)
	fmt.Printf("%.15f\n", b)

	// 通过自动推导创建的浮点数类型是float64
	c := 123.456

	fmt.Printf("%T\n", c)
}

func main() {
	price := 3.2
	var weight float64
	fmt.Scan(&weight)

	sum := price * weight

	fmt.Printf("%.2f\n", sum)
}
