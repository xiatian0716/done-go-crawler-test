package engine

import (
	"go-crawler-test/fetcher"
	"log"
)

// 将worker独立处理
// fetcher网页body(Url)
func worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching url:%s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch Error "+
			"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	// parse网页item(ParseFunc)
	return r.ParseFunc(body), nil
}
