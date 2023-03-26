/*
 * @Author MierX
 * @Title 01error接口的使用
 * @Date 2023-03-26 周一
 * @Time 11:38:31
 */

package main

import (
	"errors"
	"fmt"
)

func main() {
	//var err1 = fmt.Errorf("%s", "this is normol err")
	err1 := fmt.Errorf("%s", "this is normal err")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is normal err2")
	fmt.Println("err2 = ", err2)
}
