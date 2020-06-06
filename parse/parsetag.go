package parse

import (
	"go-crawler-test/engine"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func ParseTag(content []byte) engine.ParseResult {
	//<a href="/tag/科普" class="tag">科普</a>
	re := regexp.MustCompile(regexpStr)

	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requesrts = append(result.Requesrts, engine.Request{
			Url:       "https://book.douban.com" + string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}

	return result
}
