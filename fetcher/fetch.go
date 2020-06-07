package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 请求频率控制(100毫秒-1秒10个请求)
var rateLimiter = time.Tick(100 * time.Millisecond) // 100毫秒
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	// 设置请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: get url:%s", url)
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
	return ioutil.ReadAll(resp.Body)
}
