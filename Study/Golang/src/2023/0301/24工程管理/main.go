/*
 * @Author MierX
 * @Title main
 * @Date 2023-03-13 周一
 * @Time 14:48:55
 */

package main

import (
	"2023/0301/24工程管理/calc"
	"fmt"
)

func init() {
	fmt.Println("init")
}

func main() {
	rs := calc.Add(1, 3)
	fmt.Println(rs)
}
