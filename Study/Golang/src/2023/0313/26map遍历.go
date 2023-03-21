/*
 * @Author MierX
 * @Title 26map遍历
 * @Date 2023-03-21 周一
 * @Time 14:51:45
 */

package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "yo", 3: "go"}

	//第一个返回值为key，第二个返回值为value，遍历结果是无序的
	for i, s := range m {
		fmt.Printf("%d ====> %s\n", i, s)
	}

	//如果判断一个key值是否存在
	//第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在为真
	value, ok := m[0]
	if ok == true {
		fmt.Println("m[0] = ", value)
	} else {
		fmt.Println("不存在")
	}
}
