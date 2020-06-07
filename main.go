package main

import (
	"go-crawler-test/engine"
	"go-crawler-test/parse"
	"go-crawler-test/scheduler"
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

	// // 完成单任务爬虫
	// engine.Run(engine.Request{
	// 	Url:       "https://book.douban.com",
	// 	ParseFunc: parse.ParseTag,
	// })

	// 完成单任务爬虫
	// e:= 因为是个指针接受者我们要定义一个变量
	e:=engine.ConcurrentEngine{
		// & 取地址，因为SimpleScheduler是个*接受者
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})
}
