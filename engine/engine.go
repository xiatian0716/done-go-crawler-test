package engine

import (
	"go-crawler-test/fetcher"
	"log"
)

func Run(seeds ...Request) {
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
		log.Printf("Fetching url:%s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch Error "+
				"fetching url %s: %v", r.Url, err)
			continue
		}

		// parse网页item(ParseFunc)
		parseresult := r.ParseFunc(body)
		requests = append(requests, parseresult.Requesrts...) // 装载Requesrts
		for _, item := range parseresult.Items {              // 打印Items
			log.Printf("Got item:%s", item)
		}
	}
}
