/*
 * @Author MierX
 * @Title 13json解析到结构体
 * @Date 2023-03-27 周一
 * @Time 15:10:59
 */

package main

import (
	"encoding/json"
	"fmt"
)

// IT01 成员变量名首字母必须大写
type IT01 struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:",string"`
	Price    float64  `json:"price"`
}

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

	var tmp IT01 //定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Printf("tmp = %+v\n", tmp)
	}
}
