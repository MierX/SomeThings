/*
 * @Author MierX
 * @Title 12通过map生成json
 * @Date 2023-03-27 周一
 * @Time 15:06:34
 */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"GO", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666

	//编码成json
	rs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("rs = ", string(rs))
	}
}
