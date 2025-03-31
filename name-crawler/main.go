package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 创建输出文件
	file, err := os.Create("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 爬取多个页面
	for page := 1; page <= 5; page++ {
		url := fmt.Sprintf("http://www.namechina.cn/xingming/list_%d.html", page)
		names := crawlPage(url)

		// 写入文件
		for _, name := range names {
			_, err := file.WriteString(name + "\n")
			if err != nil {
				log.Printf("写入失败: %v", err)
			}
		}
	}

	fmt.Println("爬取完成，结果已保存到 names.txt")
}

func crawlPage(url string) []string {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("创建请求失败: %v", err)
		return nil
	}

	// 设置 User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求失败: %v", err)
		return nil
	}
	defer resp.Body.Close()

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("解析HTML失败: %v", err)
		return nil
	}

	var names []string

	// 根据网站的 HTML 结构选择合适的选择器
	doc.Find(".name-list li").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		if name != "" {
			names = append(names, name)
		}
	})

	return names
}
