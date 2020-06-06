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
	Items     []interface{}
}
