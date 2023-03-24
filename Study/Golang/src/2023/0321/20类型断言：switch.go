/*
 * @Author MierX
 * @Title 20类型断言：switch
 * @Date 2023-03-24 周一
 * @Time 17:57:22
 */

package main

import "fmt"

type Student11 struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello"
	i[2] = Student11{"mike", 666}

	//类型查询，类型断言
	//第一个返回下表，第二个返回下标对应的值

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为 int，内容为 %d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为 string，内容为 %s\n", index, value)
		case Student11:
			fmt.Printf("x[%d] 类型为 Student11，内容为 %v\n", index, value)
		}
	}
}
