/*
 * @Author MierX
 * @Title 02不定参函数.go
 * @Date 2023-01-18 周一
 * @Time 16:28:18
 */

package main

import "fmt"

// ...不定参 在函数调用时可以传递不定量（0~n）的参数
// 不定参使用的数据格式是切片，不同于数组
func sum(arr ...int) {
	// arr 是一个数据的集合
	// fmt.Println(arr)

	//count := len(arr)
	sum := 0
	// for i := 0; i < count; i++ {
	// 	sum += arr[i]
	// }

	// 通过 for range 快速遍历集合的数据
	// 第一个变量接收集合的下标 第二个变量接收对应下标的值 若不需要使用可以用匿名变量代替
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)

}

func main() {
	sum(1, 2, 3, 4, 5)
}
