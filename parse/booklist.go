package parse

import (
	"go-crawler-test/engine"
	"regexp"
)

// <a href="https://book.douban.com/subject/35044046/" title="神经网络与深度学习"
const BooklistRe = `<a.*?href="([^"]+)" title="([^"]+)"`

func ParseBookList(contents []byte) engine.ParseResult {

	//fmt.Printf("%s",contents)

	re := regexp.MustCompile(BooklistRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requesrts = append(result.Requesrts, engine.Request{
			Url:       string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}

	return result
}
