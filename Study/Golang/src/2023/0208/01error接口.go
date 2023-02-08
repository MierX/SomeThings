/*
 * @Author MierX
 * @Title 01error接口
 * @Date 2023-02-08 周一
 * @Time 14:35:46
 */

package main

import (
	"errors"
	"fmt"
)

// panic 表示 异常
// runtime error 表示 运行时错误

func dive(a, b int) (value int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	value = a / b
	return
}

func main() {
	a := 10
	b := 0
	value, err := dive(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
