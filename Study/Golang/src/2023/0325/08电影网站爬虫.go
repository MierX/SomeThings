/*
 * @Author MierX
 * @Title 08电影网站爬虫
 * @Date 2023-03-30 周四
 * @Time 15:58:19
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func MovieToFile(i int, fileTitle, fileContent []string) {
	movie, err1 := os.Create("./Study/Golang/src/2023/0325/Spider/电影网第" + strconv.Itoa(i) + "页.txt")
	if err1 != nil {
		fmt.Println("os.Create err =", err1)
		return
	}

	defer movie.Close()

	for index, title := range fileTitle {
		movie.WriteString("【" + title + "】\n" + fileContent[index] + "\n======================================\n")
	}
}

func SpiderMovie(url string) (title, introduce string, errMsg error) {
	//开始爬取内容（将所有的网站的内容全部爬下来）
	content, err := HttpGet01(url)
	if err != nil {
		fmt.Println("HttpGet err =", err)
		errMsg = err
		return
	}

	//解析表达式 <h2 data-v-63864230="" class="m-b-sm">霸王别姬 - Farewell My Concubine</h2>
	reg := regexp.MustCompile(`<h2 data-v-63864230="" class="m-b-sm">(?s:(.*?)) - `)
	if reg == nil {
		errMsg = fmt.Errorf("regexp.MustCompile err")
		return
	}

	//取出标题
	titleBuf := reg.FindAllStringSubmatch(content, 1) //只过滤第一个
	for _, titleData := range titleBuf {
		title = titleData[1]
	}

	//解析表达式 <p data-v-63864230=""></p>
	reg = regexp.MustCompile(`<p data-v-63864230="">(?s:(.*?))</p>`)
	if reg == nil {
		errMsg = fmt.Errorf("regexp.MustCompile err")
		return
	}

	//取出间接
	introduceBuf := reg.FindAllStringSubmatch(content, 1) //只过滤第一个
	for _, introduceData := range introduceBuf {
		introduce = introduceData[1]
		introduce = strings.Replace(introduce, "\r\n", "", -1)
		introduce = strings.Replace(introduce, "\r", "", -1)
		introduce = strings.Replace(introduce, "\n", "", -1)
		introduce = strings.Replace(introduce, "\t", "", -1)
		introduce = strings.Replace(introduce, " ", "", -1)
		introduce = strings.Replace(introduce, "	", "", -1)
		introduce = strings.Replace(introduce, "。", "。\n", -1)
	}

	return
}

// HttpGet01 爬取网页内容
func HttpGet01(url string) (result string, errMsg error) {
	resp, err := http.Get(url)
	if err != nil {
		errMsg = err
		return
	}

	defer resp.Body.Close()

	//读取网页body内容
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println("resp.Body.Read err =", err)
			break
		}

		result += string(buf[:n])
	}

	return
}

// SpiderPage01 并发爬取
func SpiderPage01(i int, page chan<- int) {
	//明确爬取的url
	url := "https://ssr1.scrape.center/page/" + strconv.Itoa(i)
	fmt.Printf("正在爬取第%d页内容：%s\n", i, url)

	//开始爬取内容（将所有的网站的内容全部爬下来）
	content, err := HttpGet01(url)
	if err != nil {
		fmt.Println("HttpGet err =", err)
		return
	}

	//解析表达式 <a data-v-7f856186="" href="/detail/
	reg := regexp.MustCompile(`<a data-v-7f856186="" href="(?s:(.*?))"`)
	if reg == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	//利用正则表达式过滤内容
	movieUrls := reg.FindAllStringSubmatch(content, -1)

	fileTitle, fileContent := make([]string, 0), make([]string, 0)

	//从切片中取出网址
	for _, movieUrl := range movieUrls {
		movieTitle, movieIntroduce, err := SpiderMovie("https://ssr1.scrape.center" + movieUrl[1])
		if err != nil {
			fmt.Println("SpiderMovie err =", err)
			continue
		}

		fileTitle = append(fileTitle, movieTitle)
		fileContent = append(fileContent, movieIntroduce)
	}

	//把内容写入到文件
	MovieToFile(i, fileTitle, fileContent)

	page <- i

	return
}

// DoWork01 爬虫方法
func DoWork01(start, end int) {
	page := make(chan int)

	fmt.Printf("准备爬取第%d页到第%d页的网址\n", start, end)

	for i := start; i <= end; i++ {
		//定义一个函数，爬主页面
		go SpiderPage01(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页爬取完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>=起始页）：")
	fmt.Scan(&end)

	DoWork01(start, end)
}
