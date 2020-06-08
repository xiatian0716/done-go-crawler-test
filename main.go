package main

import (
	"go-crawler-test/engine"
	"go-crawler-test/parse"
	"go-crawler-test/persist"
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

	// // Scheduler实现I-所有Worker公用一个输入
	// // Scheduler实现II-并发分发Request
	// // e:= 因为是个指针接受者我们要定义一个变量
	// e := engine.ConcurrentEngine{
	// 	// & 取地址，因为SimpleScheduler是个*接受者
	// 	Scheduler:   &scheduler.SimpleScheduler{},
	// 	WorkerCount: 10,
	// }
	// e.Run(engine.Request{
	// 	Url:       "https://book.douban.com",
	// 	ParseFunc: parse.ParseTag,
	// })

	// // Scheduler实现III-Request队列和Worker队列
	// // e:= 因为是个指针接受者我们要定义一个变量
	// e := engine.ConcurrentEngine{
	// 	// & 取地址，因为SimpleScheduler是个*接受者
	// 	Scheduler:   &scheduler.QueuedScheduler{},
	// 	WorkerCount: 10,
	// }
	// e.Run(engine.Request{
	// 	Url:       "https://book.douban.com",
	// 	ParseFunc: parse.ParseTag,
	// })

	// // Scheduler重构
	// // e:= 因为是个指针接受者我们要定义一个变量
	// e := engine.ConcurrentEngine{
	// 	// & 取地址，因为SimpleScheduler是个*接受者
	// 	Scheduler: &scheduler.SimpleScheduler{},
	// 	// Scheduler:   &scheduler.QueuedScheduler{},
	// 	WorkerCount: 10,
	// }
	// e.Run(engine.Request{
	// 	Url:       "https://book.douban.com",
	// 	ParseFunc: parse.ParseTag,
	// })

	// ItemSaver的架构
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})
}
