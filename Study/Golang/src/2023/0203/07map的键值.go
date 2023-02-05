/*
 * @Author MierX
 * @Title 07map的值
 * @Date 2023-02-05 周日
 * @Time 14:39:10
 */

package main

import "fmt"

func main0701() {
	// map中的key类型必须支持 == !=，必须有大小
	m := make(map[[2]int]int)
	fmt.Println(m)

	arr := [2]int{1, 2}
	m[arr] = 101
	fmt.Println(arr)

	m1 := make(map[string]map[string]int)
	fmt.Println(m1)
	m1["小明"] = map[string]int{"语文": 100, "数学": 95}
	fmt.Println(m1)
	fmt.Println(m1["小明"])
	fmt.Println(m1["小明"]["数学"])
}

func main0702() {
	m := make(map[int]string)
	m[101] = "张三"

	// 可以通过这种方式判断map中是否存在某个键
	v, ok := m[102]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("未找到数据")
	}
}

func main() {
	m := make(map[int]string)
	m[101] = "张三"
	fmt.Println(m)
	// 删除mqp中指定key的一个元素
	delete(m, 101)
	fmt.Println(m)
}
