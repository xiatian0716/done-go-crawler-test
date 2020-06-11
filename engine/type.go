package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}

type ParseResult struct {
	Requesrts []Request
	Items     []Item
}

type Item struct {
	// 添加URL-Type-ID
	Url  string // 任何爬虫都需要Url值
	Type string // elasticsearch数据table名
	Id   string // 在数据存储层面去重复-id行

	Payload interface{} //Payload你随便放什么值都可以Url-Id我是一定要要的
}
