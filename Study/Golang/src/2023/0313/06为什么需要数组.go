/*
 * @Author MierX
 * @Title 06为什么需要数组
 * @Date 2023-03-13 周一
 * @Time 15:34:16
 */

package main

import "fmt"

func main() {
	//id1 := 1
	//id2 := 2
	//id3 := 3

	//数组，同一个类型的集合
	var id [50]int

	//操作数据，通过下标，从0开始，到len()-1
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}

	fmt.Println("id = ", id)
}
