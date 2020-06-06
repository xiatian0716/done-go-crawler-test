package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func main() {
	// 设置请求
	req, err := http.NewRequest("GET", "https://book.douban.com/", nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
	client := &http.Client{Timeout: time.Second * 10}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	// 打印StatusCode
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
	}

	// 解析body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	// fmt.Printf("%s\n", string(body))
	parseContent(body)
}

// 4-消除噪音正则表达式获取信息
func parseContent(content []byte) {
	// ()   分组用
	// +    至少一个或多个
	// [^"] 不包含"这个字符
	//<a href="/tag/科普" class="tag">科普</a>
	re := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^<]+)</a>`)

	matches := re.FindAllSubmatch(content, -1)

	for _, m := range matches {
		fmt.Printf("url:%s\n", "https://book.douban.com"+string(m[1]))
	}
}
