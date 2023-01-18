/*
 * @Author MierX
 * @Title 08百钱百鸡.go
 * @Date 2023-01-18 周一
 * @Time 13:58:04
 */

package main

import "fmt"

func main() {
	money := 100
	cockPrice := 5
	henPrice := 3
	chickenPrice := 1
	plan := 0
	count := 0
	for cock := 0; cock <= money/cockPrice; cock++ {
		for hen := 0; hen <= int(money/henPrice); hen++ {
			chicken := 100 - cock - hen
			count++
			if cock*cockPrice+hen*henPrice+chicken*chickenPrice/3 == money && cock+hen+chicken == 100 {
				plan++
				fmt.Printf("方案%d，公鸡：%d只%d元，母鸡：%d只%d元，小鸡：%d只%d元，总计：%d元\n", plan, cock, cock*cockPrice, hen, hen*henPrice, chicken, chicken*chickenPrice/3, cock*cockPrice+hen*henPrice+chicken*chickenPrice/3)
			}
		}
	}

	fmt.Printf("累计循环%d次", count)
}
