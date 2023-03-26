/*
 * @Author MierX
 * @Title 08正则表达式1
 * @Date 2023-03-26 周一
 * @Time 16:59:31
 */

package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"
	//1、解释规则，它会解析正则表达式，如果成功返回解释器
	//reg := regexp.MustCompile(`a.c`) //首尾必须是ac，中间任意
	reg := regexp.MustCompile(`a\dc`) //首尾必须ac，中间必须是数字
	if reg == nil {
		fmt.Println("err = ", reg)
		return
	}
	//2、根据规则提取关键信息
	rs := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("rs = ", rs)
}
