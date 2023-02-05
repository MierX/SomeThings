/*
 * @Author MierX
 * @Title 10结构体数组
 * @Date 2023-02-05 周日
 * @Time 15:39:14
 */

package main

import "fmt"

type Student struct {
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main1001() {
	// 结构体数组
	var arr = [5]Student{
		{1, "张三", 18, "男", 100, "中国"},
		{2, "李四", 20, "男", 95, "印度"},
		{3, "王五", 17, "女", 99, "美国"},
		{4, "赵六", 18, "男", 77, "日本"},
		{5, "宋七", 19, "女", 57, "韩国"},
	}
	fmt.Println(arr)

	arr[1].score = 45
	fmt.Println(arr)

	arr = Sort(arr)
	fmt.Println(arr)
}

// 结构体数组作为函数参数

func Sort(arr [5]Student) [5]Student {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j].age > arr[j+1].age {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	// 结构体切片
	var slice = []Student{
		{1, "张三", 18, "男", 100, "中国"},
		{2, "李四", 20, "男", 95, "印度"},
		{3, "王五", 17, "女", 99, "美国"},
		{4, "赵六", 18, "男", 77, "日本"},
		{5, "宋七", 19, "女", 57, "韩国"},
	}
	fmt.Println(slice)
	slice = append(slice, Student{6, "魏八", 23, "男", 33, "加拿大"})
	fmt.Println(slice)
	Sort1(slice)
	fmt.Println(slice)
}

// 结构体切片作为函数参数

func Sort1(slice []Student) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j].age < slice[j+1].age {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
