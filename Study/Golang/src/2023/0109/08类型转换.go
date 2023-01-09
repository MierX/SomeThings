/*
 * @Author MierX
 * @Title 08类型转换.go
 * @Date 2023-01-09 周一
 * @Time 14:30:48
 */

package main

import "fmt"

func main010901() {
	a, b, c := 0, 0, 0
	fmt.Scan(&a, &b, &c)

	sum := a + b + c
	fmt.Println(sum)
	fmt.Println(sum / 3)

	// 类型转换 数据类型（变量） 数据类型（表达式）
	fmt.Printf("%.2f\n", float64(sum)/3)
}

func main010902() {
	var d float32 = 12.34
	var e float64 = 22.22
	fmt.Printf("%.2f\n", float64(d)+e)
}

func main010903() {
	var a float64 = 3.999

	b := int(a)

	fmt.Println(b)
}

func main() {
	s := 107653
	fmt.Println("天：", s/60/60/24%365)
	fmt.Println("时：", s/60/60%24)
	fmt.Println("分：", s/60%60)
	fmt.Println("秒：", s%60)
}
