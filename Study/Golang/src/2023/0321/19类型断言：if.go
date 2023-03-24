/*
 * @Author MierX
 * @Title 19类型断言：if
 * @Date 2023-03-24 周一
 * @Time 17:40:59
 */

package main

import "fmt"

type Student10 struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello"
	i[2] = Student10{"mike", 666}

	//类型查询，类型断言
	//第一个返回下表，第二个返回下标对应的值
	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为 int，内容为 %d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为 string，内容为 %s\n", index, value)
		} else if value, ok := data.(Student10); ok == true {
			fmt.Printf("x[%d] 类型为 Student10，内容为 %v\n", index, value)
		}
	}
}
