/*
 * @Author MierX
 * @Title 06百度贴吧小爬虫
 * @Date 2023-03-30 周四
 * @Time 14:35:08
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// HttpGet 爬取网页内容
func HttpGet(url string) (result string, errMsg error) {
	resp, err := http.Get(url)
	if err != nil {
		errMsg = err
		return
	}

	defer resp.Body.Close()

	//读取网页body内容
	buf := make([]byte, 4*1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err =", err)
			break
		}

		result += string(buf[:n])
	}

	return
}

// DoWork 爬虫方法
func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//明确目标（要知道你准备在哪个范围或者网站去搜索）
	//https://tieba.baidu.com/f?kw=golang&fr=ala0&tpl=5&dyTabStr=MCw2LDEsNSwzLDQsNyw4LDIsOQ%3D%3D
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		//fmt.Println("url =", url)

		//爬取内容（将所有的网站的内容全部爬下来）
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err =", err)
			continue
		}

		//把内容写入到文件
		html, err1 := os.Create("./Study/Golang/src/2023/0325/" + strconv.Itoa(i) + ".html")
		if err1 != nil {
			fmt.Println("os.Create err =", err1)
			continue
		}

		html.WriteString(result)

		html.Close()
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>=起始页）：")
	fmt.Scan(&end)

	DoWork(start, end)
}
