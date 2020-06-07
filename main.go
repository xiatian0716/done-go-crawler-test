package main

import (
	"go-crawler-test/engine"
	"go-crawler-test/parse"
)

func main() {
	// // parse.ParseTag
	// engine.Run(engine.Request{
	// 	Url:       "https://book.douban.com",
	// 	ParseFunc: parse.ParseTag,
	// })

	// // parse.ParseBookList
	// engine.Run(engine.Request{
	// 	Url:       "https://book.douban.com/tag/神经网络",
	// 	ParseFunc: parse.ParseBookList,
	// })

	// // parse.ParseBookDetail
	// engine.Run(engine.Request{
	// 	Url:       "https://book.douban.com/subject/30293801/",
	// 	ParseFunc: parse.ParseBookDetail,
	// })

	// 完成单任务爬虫
	engine.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})
}
