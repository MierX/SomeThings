/*
 * @Author MierX
 * @Title 17切片的截取
 * @Date 2023-03-13 周一
 * @Time 17:25:22
 */

package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max] 取下标从low开始的元素，len=high-low cap=max-low
	s1 := array[:] //[0:len(array):len(array)] 不指定，长度和容量一样
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))

	//操作某个元素，和数组操作方式一样
	data := array[1]
	fmt.Println("data = ", data)

	s2 := array[3:6:7]
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	s3 := array[:6] //从0开始，去6个元素，容量也是6，常用
	fmt.Printf("s3 = %v, len = %d, cap = %d\n", s3, len(s3), cap(s3))
}
