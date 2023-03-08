/*
 * @Author MierX
 * @Title 10数字累加
 * @Date 2023-03-08 周一
 * @Time 16:15:06
 */

package main

import "fmt"

func MyFunc16() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func MyFunc17(a int) int {
	if a == 1 {
		return a
	}
	return a + MyFunc17(a-1)
}

func main() {
	fmt.Println(MyFunc16())
	fmt.Println(MyFunc17(100))
}
