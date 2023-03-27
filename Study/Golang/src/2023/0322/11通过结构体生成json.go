/*
 * @Author MierX
 * @Title 11通过结构体生成json
 * @Date 2023-03-27 周一
 * @Time 14:11:01
 */

package main

import (
	"encoding/json"
	"fmt"
)

// IT 成员变量名首字母必须大写
type IT struct {
	Company  string   `json:"company"` //可以将json键名首字母转小写
	Subjects []string `json:"-"`       //隐藏此字段
	IsOk     bool     `json:",string"` //将此字段转换成字符串
	Price    float64
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"GO", "C++", "Python", "Test"}, true, 666.666}

	//编码，根据内容生成json文本
	rs, err := json.Marshal(s) //返回的切片
	if err != nil {
		fmt.Println("err = ", err)
		return
	} else {
		fmt.Println("rs = ", string(rs))
	}
}
