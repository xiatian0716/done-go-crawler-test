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

	// parse.ParseBookList
	engine.Run(engine.Request{
		Url:       "https://book.douban.com/tag/神经网络",
		ParseFunc: parse.ParseBookList,
	})
}
