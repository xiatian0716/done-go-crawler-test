package main

import (
	"go-crawler-test/engine"
	"go-crawler-test/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseContent,
	})
}
