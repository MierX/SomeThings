/*
 * @Author MierX
 * @Title 09正则表达式2
 * @Date 2023-03-26 周一
 * @Time 17:09:40
 */

package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 ab 11.23 7. 8.99 abc 6.66 6.67a a1.22"
	//1、解释规则，它会解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`\d+\.\d+`) //找出有效小数
	if reg == nil {
		fmt.Println("err = ", reg)
		return
	}
	//2、根据规则提取关键信息
	//rs := reg.FindAllStringSubmatch(buf, -1)
	rs := reg.FindAllString(buf, -1)
	fmt.Println("rs = ", rs)
}
