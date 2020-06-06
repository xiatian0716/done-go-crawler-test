package engine

import (
	"go-crawler-test/fetcher"
	"log"
)

func Run(seeds ...Request) {

	var requests []Request

	for _, e := range seeds {
		requests = append(requests, e)
	}
	// log.Print(requests)

	for len(requests) > 0 {

		r := requests[0]

		// 爬取URL
		requests = requests[1:]
		log.Printf("Fetching url:%s", r.Url)
		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("Fetch Error: %s", r.Url)
		}

		parseresult := r.ParseFunc(body)

		requests = append(requests, parseresult.Requesrts...)
		// log.Print("\n", requests, "\n")

		for _, item := range parseresult.Items {
			log.Printf("Got item:%s", item)
		}
	}
}
