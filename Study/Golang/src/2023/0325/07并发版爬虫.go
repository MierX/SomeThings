/*
 * @Author MierX
 * @Title 07并发版爬虫
 * @Date 2023-03-30 周四
 * @Time 14:55:54
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

// SpiderPage 并发爬取
func SpiderPage(i int, page chan<- int) {
	url := "https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬取第%d页网页：%s\n", i, url)

	//爬取内容（将所有的网站的内容全部爬下来）
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err =", err)
		return
	}

	//把内容写入到文件
	html, err1 := os.Create("./Study/Golang/src/2023/0325/" + strconv.Itoa(i) + ".html")
	if err1 != nil {
		fmt.Println("os.Create err =", err1)
		return
	}

	html.WriteString(result)

	html.Close()

	page <- i

	return
}

// DoWork 爬虫方法
func DoWork(start, end int) {
	page := make(chan int)
	//明确目标（要知道你准备在哪个范围或者网站去搜索）
	//https://tieba.baidu.com/f?kw=golang&fr=ala0&tpl=5&dyTabStr=MCw2LDEsNSwzLDQsNyw4LDIsOQ%3D%3D
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
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
