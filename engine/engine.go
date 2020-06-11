package engine

import (
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
