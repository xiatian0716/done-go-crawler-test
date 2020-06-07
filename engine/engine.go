package engine

import (
	"go-crawler-test/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	// 把seeds放到Q队列
	var requests []Request
	for _, e := range seeds {
		requests = append(requests, e)
	}

	for len(requests) > 0 {
		// 取出Request(Url+ParseFunc)
		r := requests[0]
		requests = requests[1:]

		// fetcher网页body(Url)
		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		// parse网页item(ParseFunc)
		requests = append(requests, parseResult.Requesrts...) // 装载Requesrts
		for _, item := range parseResult.Items {              // 打印Items
			log.Printf("Got item:%s", item)
		}
	}
}

// fetcher网页body(Url)
func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching url:%s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch Error "+
			"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	// parse网页item(ParseFunc)
	return r.ParseFunc(body), nil
}
