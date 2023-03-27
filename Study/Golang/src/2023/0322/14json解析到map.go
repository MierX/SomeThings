/*
 * @Author MierX
 * @Title 14json解析到map
 * @Date 2023-03-27 周一
 * @Time 15:22:01
 */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonBuf := `
		{
			"company": "itcast",
			"subjects": [
				"GO",
				"C++",
				"Python",
				"Test"
			],
			"IsOk": "true",
			"price": 666.666
		}
		`

	//创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Printf("m = %+v\n", m)
	}

	//var str string
	//不能直接取出来用
	//str = m["company"]

	//类型断言
	for s, i := range m {
		//fmt.Printf("%v =========> %v\n", s, i)
		switch data := i.(type) {
		case string:
			fmt.Printf("m[%s]的值类型为string，内容为%v\n", s, data)
		case []interface{}:
			fmt.Printf("m[%s]的值类型为[]interface{}，内容为%v\n", s, data)
		case bool:
			fmt.Printf("m[%s]的值类型为bool，内容为%v\n", s, data)
		case float64:
			fmt.Printf("m[%s]的值类型为float64，内容为%v\n", s, data)
		}
	}
}
