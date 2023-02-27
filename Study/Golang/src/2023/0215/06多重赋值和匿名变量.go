/*
 * @Author MierX
 * @Title 06多重赋值和匿名变量
 * @Date 2023-02-27 周一
 * @Time 15:26:02
 */

package main

import "fmt"

func main() {
	//a := 10
	//b := 20
	//c := 30

	a, b := 10, 20

	//交换两个变量的值
	tmp := a
	a = b
	b = tmp
	fmt.Printf("a = %d, b = %d\n", a, b)

	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	i, j = 10, 20
	//_匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用，才有又是
	tmp, _ = i, j
	fmt.Println("tmp = ", tmp)

	var c, d, e int
	c, d, e = test() // return 123
	fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)

	_, d, _ = test()
	fmt.Println("d = ", d)
}

func test() (a, b, c int) {
	return 1, 2, 3
}
