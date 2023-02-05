/*
 * @Author MierX
 * @Title 06map练习题
 * @Date 2023-02-05 周日
 * @Time 14:26:37
 */

package main

import "fmt"

func main() {
	// 通过键盘输入一个英文字符串 统计每个字母出现的次数
	var str string
	fmt.Scan(&str)
	fmt.Println(str)
	slice := []byte(str)
	m := make(map[string]int)
	for i := 0; i < len(slice); i++ {
		m[string(slice[i])]++
	}
	fmt.Println(m)
}
